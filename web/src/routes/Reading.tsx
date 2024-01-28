import React from 'react';
import {AppLayout} from '@/components/layouts/AppLayout';
import {Button} from "@/components/ui/button";

export const Reading: React.FC = () => {
    let mediaRecorder, chunks = [], audioURL = ''

    // mediaRecorder setup for audio
    if(navigator.mediaDevices && navigator.mediaDevices.getUserMedia){
        console.log('mediaDevices supported..')
        navigator.mediaDevices.getUserMedia({
            audio: true
        }).then(stream => {
            mediaRecorder = new MediaRecorder(stream)
            mediaRecorder.ondataavailable = (e) => {
                chunks.push(e.data)
                console.log(e.data);
            }
            mediaRecorder.onstop = () => {
                const blob = new Blob(chunks, {'type': 'audio/ogg; codecs=opus'})
                chunks = []
                audioURL = window.URL.createObjectURL(blob)
                // document.querySelector('audio').src = audioURL
            }
            const myStream = mediaRecorder.stream;
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

    return (
        <AppLayout>
            <Button onClick={record}>Record</Button>
            <Button onClick={stopRecording}>Stop</Button>
            <Button onClick={downloadAudio}>Download</Button>
        </AppLayout>
    );
};
