import React, {useEffect, useState} from 'react';
import {AppLayout} from '@/components/layouts/AppLayout';
import {Button} from "@/components/ui/button";
import {APIAudio} from "../lib/api/APIAudio.tsx";

export const Reading: React.FC = () => {
    const [mediaRecorder, setMediaRecorder] = useState<MediaRecorder>(new MediaRecorder(new MediaStream()));
    const chunks: Array<Blob> = [];
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
                const myStream = rec.stream;
                console.log(myStream);
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

    return (
        <AppLayout>
            <Button onClick={record}>Record</Button>
            <Button onClick={stopRecording}>Stop</Button>
            <Button onClick={downloadAudio}>Download</Button>
            <Button onClick={uploadAudio}>Upload</Button>
        </AppLayout>
    );
};
