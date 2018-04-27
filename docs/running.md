## Running pac-mule

To start a basic instance of pac-mule you can fill in the needed paramters and run this command:

```
MULE_PAC_FILE=<PAC FILE> ./pac-mule start
```

Now clients can get the config from `http://localhost:8080/config`

You can also launch a more secure instance of pac-mule by generating a SSL certificate and key. Doing this you can run this command:

```
MULE_PAC_FILE=<PAC FILE> MULE_TLS=true MULE_PORT=443 MULE_CERT=<SSL CERT FILE> MULE_KEY=<SSL KEY FILE> ./pac-mule start
```

Now clients can get the config from `https://localhost/config`
