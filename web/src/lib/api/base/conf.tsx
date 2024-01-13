export const devMode = window.location.host.includes('localhost');
const devhost = 'http://localhost:3000';
const prodAPIHost = window.location.origin;
export const url = devMode ? devhost : prodAPIHost;
export const APIConfig = {
    timeout: 30000,
    baseURL: url + "/api/",
    contentType: "application/json",
    headers: {
        "Accept": "application/json",
        "Accept-Language": "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7",
        "Pragma": "no-cache",
        "Cache-Control": "no-cache",
        "Content-Type": "application/json;charset=UTF-8",
    },
};
