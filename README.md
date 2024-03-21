# Invictus Radio

Frontend voor de radio in het Colosseum, gebaseerd op N.A.U.T.
Is in principe gewoon een frontend voor spotify.

## Hoe te ontwikkelen

Voordat je kan ontwikkelen heb je een aantal dingen nodig, namelijk:

- Node 18.19.0 of hoger
- Golang 1.21 of hoger

Daarna kan je de dependencies installeren, eerst voor de frontend met

```console
npm i
```

en voor de backend met

```console
cd scheduler && go get
```

De devserver moet je apart installeren, dat doe je met

```console
go install github.com/cosmtrek/air@latest
```

De backend maakt gebruikt van GNU Make, vogel vooral uit voor je eigen disto hoe je dit installeerd

```console
sudo pacman -S make
```

_vaak wordt dit meegeleverd op Linux_

### Draaien

Frontend:

```console
npm run dev
```

Backend:

```console
cd scheduler; make dev
```

Voor styling gebruik ik PicoCSS v2. Documentatie is [hier](https://v2.picocss.com/docs) te vinden
