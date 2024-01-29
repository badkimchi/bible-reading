import React, {useState} from 'react';
import {AppLayout} from '@/components/layouts/AppLayout';
import {Button} from "@/components/ui/button";
import {APIAudio} from "../lib/api/APIAudio.tsx";

export const Reading: React.FC = () => {
    const [mediaRecorder, setMediaRecorder] = useState<MediaRecorder>(new MediaRecorder(new MediaStream()));
    let chunks: Array<Blob> = []
    let audioURL = ''

    // mediaRecorder setup for audio
    if(navigator.mediaDevices && navigator.mediaDevices.getUserMedia){
        console.log('mediaDevices supported..')
        navigator.mediaDevices.getUserMedia({
            audio: true
        }).then(stream => {
            const rec = new MediaRecorder(stream);
            setMediaRecorder(rec);
            rec.ondataavailable = (e) => {
                chunks.push(e.data)
                console.log(e.data);
            }
            rec.onstop = () => {
                const blob = new Blob(chunks, {'type': 'audio/ogg; codecs=opus'})
                chunks = []
                const url = window.URL || window.webkitURL;
                audioURL = url.createObjectURL(blob)
                const audio = document.querySelector('audio')
                if (audio) {
                    audio.src = audioURL;
                }
            }
            const myStream = rec.stream;
            console.log(myStream);
            myStream.onaddtrack = (e) => {
                console.log(e);
            }

        }).catch(error => {
            console.log('Following error has occured : ',error)
        })
    }

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
