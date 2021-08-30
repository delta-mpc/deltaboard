# Web frontend project for Deltaboard

## Local development configuration

For convenience, we use the online Deltaboard's API. If you started
a local backend server to be used by the frontend project, configure
the backend server API accordingly.

1. Copy and paste ```.env``` to ```.env.development```

2. Edit ```.env.development```:

```shell
VUE_APP_BASE_API='https://localhost:8090'
VUE_APP_VERSION=''
```

So that the API requests could be sent to vue dev server to be proxied
to support HTTPS on localhost to overcome cross-site cookie problem and
bypass cross-site origin limitations.

3. Edit ```vue.config.js```, set the target API domain to the actual domain
we will use for debug:

```json
{
  "target": "https://api.board.deltampc.com"
}
```

4. Now we can start the dev server:

```shell
$ npm run serve
```

## Standalone building and deployment

1. Edit ```.env``` file to set the backend API domain. This is usually done by
Continuous Integration system.

```shell
VUE_APP_BASE_API='https://api.board.deltampc.com'
VUE_APP_VERSION='1.0.1'
```

2. run npm command to build the project:

```shell
$ npm run build
```

3. After successful building, we have the distribution archive under ```dist```
folder. We can copy and paste the files under the folder to the target deployment
places. The deployment folder structure should be like this:

```
| - 1.0.1
|     | - css
|     | - js
|     | - ...
| - index.html
```

```index.html``` is at the root level, and all the other files are in a
sub folder whose name is the version number we set in ```.env``` file.


