# Go-Reloaded

## Description
**Go-Reloaded** est un outil en ligne de commande développé en Go. Il permet de traiter, compléter, éditer et autocorriger des fichiers texte en appliquant des règles spécifiques de formatage (conversion hexadécimale, binaire, majuscules, minuscules, capitalisation, gestion de la ponctuation et des articles a/an).

## Fonctionnalités
- Conversion de nombres hexadécimaux `(hex)` et binaires `(bin)` en décimal.
- Modification de la casse des mots : majuscules `(up)`, minuscules `(low)` et première lettre en majuscule `(cap)` avec support de la quantité (ex: `(up, 2)`).
- Correction automatique de la ponctuation (espaces, points de suspension `...`, guillemets).
- Correction des articles `a` en `an` devant les voyelles et le `h`.

## Utilisation

L'application prend en paramètre le fichier d'entrée (contenant le texte à modifier) et le fichier de sortie (où le résultat sera enregistré) :

```bash
go run . sample.txt result.txt
```
# Exemples
Hexadécimal & Binaire :

Entrée : Simply add 42 (hex) and 10 (bin)

Sortie : Simply add 66 and 2

Casse et Majuscules :

Entrée : Ready, set, go (up) !

Sortie : Ready, set, GO!

Ponctuation :

Entrée : I was sitting over there , and then BAMM !

Sortie : I was sitting over there, and then BAMM!

## Auteur
Basma Bouzadi