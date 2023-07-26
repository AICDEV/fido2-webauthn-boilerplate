# fido2 workshop angular and golang 

## overview
- [disclaimer](#disclaimer)
- [links](#links)
- [requirements](#requirements)
- [preparation](#preparation)
- [config](#config)
- [start](#start)
- [api](#api)


## disclaimer
You don't want to use this backend blindly for productive applications, please. Think about using a provider like Auth0(Okta) to prevent doing your own implementation like in this repo. [https://auth0.com/docs/secure/multi-factor-authentication/fido-authentication-with-webauthn](https://auth0.com/docs/secure/multi-factor-authentication/fido-authentication-with-webauthn)

## links

- [https://fidoalliance.org/fido2/](https://fidoalliance.org/fido2/)
- [https://webauthn.guide/](https://webauthn.guide/)
- [https://webauthn.io/](https://webauthn.io/)

## requirements

In order to follow allong with all examples you should have the following tools installed on your system:

- [docker](https://www.docker.com/) Make sure you have intalled docker with docker compose support!
- [mkcert](https://github.com/FiloSottile/mkcert) We need this tiny tool to create our local dev certificates to prevent browser issues
-  A code editor of your choice. Maybe [vim](https://www.vim.org/) ?

## preparation

Before you start you need to create your dev certificates. First of all install the local CA by running the following command:
```bash
mkcert -install 
```

After completion switch inside the ./deployment folder and run the following command:

```bash
mkcert fido.workshop "*.fido.workshop" localhost 127.0.0.1 ::1
```

You should see to following two generate *.pem files:
- fido.workshop+4-key.pem
- fido.workshop+4.pem

If, for some reasons, the filenames differ, change to *.pem files to the names above. Please restart your browser(s).

Right after we need to edit the /etc/hosts file. On windows inside the C:\Windows\System32\drivers\etc folder. Please add the following line to the hosts file:

```text
127.0.0.1 fido.workshop
```

You are ready to go!

## config
In order to make the backend work you have to create a docker volume and connect your container to the provided configuration.

1. Create local docker volume by running the following command inside the root path of this repository

```bash
mkdir -p docker_volumes/backend/config
```

2. Inside the newly created /backend/config folder run the following command to create an EC Pub-Priv key pair
```bash
# create private key
openssl ecparam -name prime256v1 -genkey -noout -out key.pem
# extract public key
openssl ec -in key.pem -pubout -out public.pem
```

## start

In order to start the local dev stack you need to run the following command:

```bash
docker compose up -d --build
```

You should see three containers up and running:

- fido_proxy
- fido_frontend
- fido_backend

You can access the frontend by entering https://fido.workshop into your browser

## api 
The backend expose the following api ressources:

- /api/v1/service/signup/begin (POST)
- /api/v1/service/signup/finish (POST)
- /api/v1/service/authenticate/begin (POST)
- /api/v1/service/authenticate/finish (POST)
- /api/v1/member/register/device/begin (POST)
- /api/v1/member/register/device/finish (POST)
- /api/v1/member/name (GET)
