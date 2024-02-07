import React, {useEffect, useState} from 'react';
import {AppLayout} from '@/components/layouts/AppLayout';
import {Button} from "@/components/ui/button";
import {APIAudio} from "../lib/api/APIAudio.tsx";
import {useLocation} from 'react-router-dom';

export const Reading: React.FC = () => {
    const location = useLocation();
    const [mediaRecorder, setMediaRecorder] = useState<MediaRecorder>(new MediaRecorder(new MediaStream()));
    const [chunks, setChunks] = useState<Array<Blob>>([]);
    const [audioURL, setAudioURL] = useState<string>('');

    useEffect(() => {
        // mediaRecorder setup for audio
        if(navigator.mediaDevices && navigator.mediaDevices.getUserMedia){
            console.log('mediaDevices supported..')
            navigator.mediaDevices.getUserMedia({
                audio: true
            }).then(stream => {
                const rec = new MediaRecorder(stream);
                setMediaRecorder(rec);
                rec.ondataavailable = (e) => {
                    chunks[0] = e.data;
                    setChunks(chunks);
                }
                rec.onstop = () => {
                    const blob = new Blob(chunks, {'type': 'audio/ogg; codecs=opus'});
                    const url = window.URL || window.webkitURL;
                    const audioURL = url.createObjectURL(blob);
                    setAudioURL(audioURL);
                    const audio = document.querySelector('audio')
                    if (audio) {
                        audio.src = audioURL;
                    }
                }
            }).catch(error => {
                console.log('Following error has occured : ',error)
            })
        }
    }, [])

    const record = () => {
        mediaRecorder.start()
    }

    const stopRecording = () => {
        mediaRecorder.stop()
    }

    const downloadAudio = () => {
        const downloadLink = document.createElement('a')
        downloadLink.href = audioURL
        downloadLink.setAttribute('download', 'audio')
        downloadLink.click()
    }

    const uploadAudio = () => {
        const blob = new Blob(chunks, {'type': 'audio/ogg; codecs=opus'});
        const formData = new FormData();
        formData.append('audioFile', blob, 'recording.ogg');
        APIAudio.postAudio(formData)
            .then(resp => console.log(resp))
            .catch(err => console.error(err));
    }
    let audioUrl = `${window.location.protocol}//${window.location.hostname}/audio/${location.pathname.split('/')[2]}`;
    if (window.location.hostname === 'localhost') {
       audioUrl = `http://localhost:3000/audio/${location.pathname.split('/')[2]}`
    }

    return (
        <AppLayout>
            <audio controls src={audioUrl}></audio>
            <Button onClick={record}>Record</Button>
            <Button onClick={stopRecording}>Stop</Button>
            <Button onClick={downloadAudio}>Download</Button>
            <Button onClick={uploadAudio}>Upload</Button>
        </AppLayout>
    );
};
