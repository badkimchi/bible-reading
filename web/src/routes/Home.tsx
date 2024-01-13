import React, {useEffect} from 'react';
import {AppLayout} from '../components/layouts/AppLayout';
import {loginInfoStore} from "../lib/stores/loginInfoStore.ts";
import {APIAccount} from "../lib/api/APIAccount.tsx";

export const Home: React.FC = () => {
    const logout = loginInfoStore(state => state.logout);
    const signOut = () => {
        logout();
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
            <>
                Home Page
                <br/>
                <br/>
                <br/>
                <button onClick={signOut}> Sign Out </button>
            </>
        </AppLayout>
    );
};
