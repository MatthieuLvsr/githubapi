# Dapp Go - Gestionnaire de Repos GitHub

Ce projet est une Dapp (Application décentralisée) en Go qui permet de gérer les dépôts GitHub d'un utilisateur. Il offre les fonctionnalités suivantes :

1. Requêter l'API GitHub pour récupérer tous les dépôts d'un utilisateur spécifié dans le fichier `.env`.
2. Cloner localement chaque dépôt, exécuter un `git pull` et un `git fetch` pour mettre à jour les dépôts.
3. Archiver tous les dépôts dans un fichier `.zip`.
4. Permettre aux utilisateurs de télécharger le fichier `.zip` associé à un nom d'utilisateur spécifié.

## Configuration

Pour exécuter cette Dapp, suivez les étapes suivantes :

1. Créez un fichier `.env` à la racine du projet avec les informations suivantes :
```bash
GITHUB_KEY=YOUR_GITHUB_API_KEY
GITHUB_USER=YOUR_TARGET
ORG_MODE=MODE_ORGANISATION (TRUE OR FALSE)
REPOS_PER_PAGE=NOMBRE_DE_REPOS_SOUHAITES
```

2. Exécutez le projet en utilisant Docker Compose :
```bash
docker-compose up
```

ou via le Makefile

```bash
make restart
make log
```
## Utilisation

Une fois la Dapp en cours d'exécution, vous pouvez accéder à l'interface utilisateur depuis votre navigateur.

Ouvrez votre navigateur et accédez à l'URL http://localhost:8081/download?name=NOM_DE_VOTRE_UTILISATEUR pour télécharger le fichier .zip associé.