# Veil of Souls

🥷 Un RPG console en Go 

---

## Description

**Veil of Souls** est un jeu de rôle (RPG) développé en Go, dans lequel le joueur crée un personnage, explore un donjon de 10 salles, combat des ennemis et améliore son équipement.  
Ce projet est basé sur le **sujet RED** et couvre l’ensemble des fonctionnalités demandées.

### Fonctionnalités principales :
- Création du personnage : nom, classe (Humain / Elfe / Nain), arme de départ  
- Inventaire limité avec possibilité d’augmentation (+10 slots, max 3 fois)  
- Objets : potions de vie / poison, ressources, équipements  
- Marchand et forgeron pour acheter / fabriquer des objets  
- Équipement (tête, torse, pieds) modifiant les PV max  
- Exploration de 10 salles avec ennemis et sortie fermée tant que tous ne sont pas battus  
- Combat tour par tour contre les monstres et un boss final  
- Système d’expérience et montée en niveau (stats augmentées)  
- Sauvegarde et chargement de la partie au format JSON  

---

## Structure du projet

veilofsouls/
│── go.mod
│── main.go
│
├── intro/ # écran d’accueil
├── character/ # gestion du personnage : stats, niveau, équipement
├── inventory/ # usage des objets dans l’inventaire
├── merchant/ # achat d’objets
├── blacksmith/ # fabrication d’équipements
├── mapgame/ # exploration des salles du donjon
├── rooms/ # définition des salles (room1.go → room10.go)
├── monster/ # définition des monstres (gobelins, boss…)
├── combat/ # logique de combat tour par tour
├── utils/ # fonctions utilitaires
└── save/ # sauvegarde et chargement du jeu


---

## Installation / Exécution

1. Installer Go (>= 1.22 recommandé)  
2. Cloner le dépôt :  

   ```bash
   git clone https://github.com/CorentinNKN/projet-red_Veil-of-Souls.git
   cd projet-red_Veil-of-Souls

## Lancer le jeu

go run main.go

## Commandes dans le jeu

Menus principaux : taper le numéro du choix

Exploration dans le donjon :

z → haut

q → gauche

s → bas

d → droite

i → accéder à l’inventaire

r → quitter la salle

Combat :

1 → attaquer

2 → lancer une boule de feu (si disponible)

3 → fuir

⚠️ On ne peut pas quitter tant que les 10 salles ne sont pas terminées.

