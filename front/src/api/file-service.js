import axios from 'axios'
import {ApiPromise} from "@/api/api-promise"


const fileService = axios.create()

fileService.interceptors.response.use(
    (response) => {

        if(response.status === 200) {
            // Normal response
            return Promise.resolve(response.data.data);

        } else if (response.status === 400) {
            // Validation error
            return Promise.reject(response.data.message);
        } else {
            // Exceptions
            return Promise.reject("server exception");
        }
    },
    (error) => {
        return Promise.reject(error);
    });

function post(url, data, config) {
    return new ApiPromise(fileService.post(url, data, config));
}

export default { fileService, post };
