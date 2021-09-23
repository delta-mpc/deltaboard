export class ApiPromise {
    constructor(promise) {
        this.userCaught = false;
        this.promise = promise;
        this.promise = this.promise.catch(this.managedOnReject.bind(this));
    }
    then(onFulfilled, onRejected) {

        if(typeof(onRejected) !== 'undefined') {
            this.userCaught = true;
        }

        this.promise = this.promise.then(onFulfilled, onRejected);
        return this;
    }
    catch(onRejected) {
        if(typeof(onRejected) !== 'undefined') {
            this.userCaught = true;
        }
        this.promise = this.promise.catch(onRejected);
        return this;
    }

    finally(onFinally) {
        this.promise = this.promise.finally(onFinally);
        return this;
    }

    managedOnReject(err) {
        if(!this.userCaught) {
            ApiPromise.defaultOnReject(err);
        }
        return Promise.reject(err);
    }
}

