# Cf-CLI-Auth
This library piggy backs off of the [Cloudflared CLI](https://github.com/cloudflare/cloudflared) to allow you to perform authentication to Cloudflare Access applications using either user based auth or service token auth.

The library implements a `http.RoundTripper` for authentication with a Cloudflare Access Application.

## Examples
The examples below show an implementation for user and service auth that return a `http.Client` for use by client wrappers.

### User auth
An example of the user auth flow 

```golang
appDomain := "https://<domain>/"
tr := http.DefaultTransport
transport, err := cf.NewAccessTokenClient(context, tr, appDomain)
if err != nil {
    return err
}
client := &http.Client{Transport: transport}
```

### Service auth
An example of the service auth flow

```golang
tr := http.DefaultTransport
transport := cf.NewServiceTokenClient(tr, clientId, clientSecret)
client := &http.Client{Transport: transport}
```

## Limitations
The library is currently limited to only perform an initial token fetch at the creation of the `RoundTripper` and as such is not designed to be used for long running operations.