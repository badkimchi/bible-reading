import {create} from 'zustand';
import {persist} from 'zustand/middleware';
import {LoginInfo} from '../models/loginInfo.ts';

type AccountState = {
    current: LoginInfo | null;
    setUser: (info: LoginInfo) => void;
    logout: () => void;
};

export const loginInfoStore = create(
    persist<AccountState>(
        (set, _) => ({
            current: new LoginInfo(),
            setUser: (info) => set({current: info}),
            logout: () => set({current: null}),
        }),
        {
            name: 'account-storage',
        }
    )
);
