
import axios from 'axios';
import {ApiPromise} from "@/api/api-promise";
import { cacheAdapterEnhancer } from 'axios-extensions';
import cache from './v1-cache.js';
let versionPrefix = "v1";
let baseUrl = process.env.VUE_APP_BASE_API + '/' + versionPrefix;
console.log(baseUrl)

let httpClient = axios.create({
    baseURL: baseUrl,
    withCredentials: true,
    adapter:cacheAdapterEnhancer(axios.defaults.adapter,{enabledByDefault:false,defaultCache:cache})
});

httpClient.interceptors.response.use(
    (response) => {
        if(response.config.responseType === 'blob') {
            return Promise.resolve(response.data);
        } else if(response.status === 200) {
            // Normal response
            if(response.data && response.data.message && response.data.message != 'success') {
               return Promise.reject(response.data.message)
            }
            return Promise.resolve(response.data.data);
            
        } else if (response.status === 400) {
            return Promise.reject(response.data);
        } else {
            // Exceptions
            return Promise.reject("server exception");
        }
    },
    (error) => {
        if(error.response.data && error.response.data.message) {
           return Promise.reject(error.response.data.message)
        }
        return Promise.reject(error);
    });
function post(url, data, config) {
    return new ApiPromise(httpClient.post(url, data, config));
}

function get(url, config) {
    return new ApiPromise(httpClient.get(url, config));
}

function put(url, data, config) {
    return new ApiPromise(httpClient.put(url, data, config));
}

function del(url, config) {
    return new ApiPromise(httpClient.delete(url, config));
}

export default { httpClient, post, get, put, del };
