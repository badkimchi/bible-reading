import React, {useEffect} from 'react';
import {AppLayout} from '@/components/layouts/AppLayout';
import {loginInfoStore} from "@/lib/stores/loginInfoStore";
import {APIAccount} from "../lib/api/APIAccount.tsx";
import {Button} from "@/components/ui/button";
import {useNavigate} from 'react-router-dom';

export const Home: React.FC = () => {
    const navigate = useNavigate();
    const logout = loginInfoStore(state => state.logout);
    const signOut = () => {
        logout();
    }
    const startReading = () => {
        navigate('/psalms/1');
    }

    useEffect(() => {
        APIAccount.getAccount()
            .then((resp) => {
                console.log(resp);
            })
            .catch(err => {
                console.error(err)
            });
    }, [])

    return (
        <AppLayout>
            <div className={'bg-blue-50 w-screen'}>
                <Button onClick={signOut}> Sign Out</Button>
            </div>
            <div  className={''}>
                <Button onClick={startReading}> Start Reading</Button>
            </div>
        </AppLayout>
    );
};
