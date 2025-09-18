# Veil of Souls

🥷 Un RPG console en Go 

---

## Description

Veil of Souls est un RPG console dans lequel le joueur :

- Crée un personnage (nom, classe : Humain / Elfe / Nain)
- Commence avec une arme de départ
- Explore un donjon de **10 salles**, chacune peuplée de monstres
- Ne peut sortir d’une salle tant que tous les ennemis ne sont pas vaincus
- Astreint à combattre un boss final à la fin du donjon
- Gagne de l’expérience, monte de niveau, améliore ses statistiques
- Dispose d’un inventaire avec capacité limitée, pouvant être augmentée (max +10 cases, jusqu’à 3 fois)
- Trouve / achète / fabrique des objets : potions, poisons, ressources, équipements
- Équipe différents types d’armures (tête, torse, pieds), qui modifient les PV maximum
- Peut interagir avec un marchand et un forgeron
- Sauvegarde et charge la partie au format JSON



##  Structure du projet

```
veilofsouls/
├── go.mod
├── main.go
├── intro/         # écran d’accueil
├── character/     # stats, niveau, équipement du personnage
├── inventory/     # gestion des objets, utilisation
├── merchant/      # achat d’objets
├── blacksmith/    # fabrication / amélioration d’équipements
├── mapgame/       # exploration des salles du donjon
├── rooms/         # définition des salles (room1.go → room10.go)
├── monster/       # définition des monstres (gobelins, boss…)
├── combat/        # logique de combat au tour par tour
├── utils/         # fonctions utilitaires
└── save/          # sauvegarde / chargement JSON
```



## 🛠 Prérequis / Installation

- Go version **1.22** ou plus recommandée
- Git (pour cloner le dépôt)



##  Lancer le jeu

```bash
git clone https://github.com/CorentinNKN/projet-red_Veil-of-Souls.git
cd projet-red_Veil-of-Souls
go run main.go
```



##  Commandes du jeu

### Menus principaux
- Entrer le numéro correspondant au choix voulu

### Exploration
- `z` : aller en haut  
- `q` : aller à gauche  
- `s` : aller en bas  
- `d` : aller à droite  
- `i` : ouvrir l’inventaire  
- `r` : quitter la salle

### Combat
- `1` : attaquer  
- `2` : lancer une boule de feu (si disponible)  
- `3` : fuir  
> ⚠️ On ne peut pas quitter tant que les 10 salles du donjon ne sont pas toutes complétées.



##  Progression et Fonctionnalités

- Système de niveaux et expérience : augmente les stats du personnage à chaque niveau
- Objets variés : consommables (potions de soins, potions de poisons), ressources, équipements
- Équipements affectant la vie maximale selon l’équipement porté
- Fabrication / amélioration via le forgeron
- Sauvegarde / chargement via fichier JSON : permet de reprendre une partie



## À venir / Améliorations possibles

- Ajouter des types de monstres plus variés
- Ajouter des compétences ou sorts supplémentaires
- Meilleure interface utilisateur (ASCII art, couleurs, etc.)
- Ajouter un système de quêtes secondaires
- Améliorer la sauvegarde pour gérer plusieurs parties
- Équilibrer la difficulté (boss, progression des ennemis, etc.)



##  Auteurs

- Corentin NOKAYA 
- Antoine MASSOUH
- Souleymane SALL 


Merci d’avoir regardé ce projet !   

