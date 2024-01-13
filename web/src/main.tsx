import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import {QueryClient, QueryClientProvider} from '@tanstack/react-query'
import {GoogleOAuthProvider} from "@react-oauth/google";

const qClient = new QueryClient();

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <GoogleOAuthProvider clientId={import.meta.env.VITE_CLIENT_ID}>
            <QueryClientProvider client={qClient}>
                <App/>
            </QueryClientProvider>
        </GoogleOAuthProvider>
    </React.StrictMode>,
)
