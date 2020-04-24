# Totenacker
Totenacker is a web service for sharing high scores and [bones files](http://www.roguebasin.com/index.php?title=Save_Files#Bones_Files)
from Roguelike games between players. Unlike [Hearse](http://hearse.krollmark.com/), this aims to offer something more
generic that could be used by any game, rather than being specifically for Nethack.

I'm not in big rush with this project as I don't really have any Roguelike games that I'm actively working on, so going
to take a very leisurely approach of building it from the domain layer outwards in an attempt at a Clean Architecture
approach in GoLang. 

## Domain Model

### Player

Someone that plays a game. The person that got the high score or produced the Bones file as a side effect of dying in
the game.

Important data points needed for the Player:
* Email: Will serve as a globally unique identifier for a player. Needs to be an email instead of a GUID, in-case we
need to communicate with the player due to abuse of the service.
* Name: A display name. Not guaranteed to be unique, how could we handle non-unique player names?

### Score

A score achieved in the game by a Player.

Important data points needed for the Score:
* Player: Who achieved the score.
* Timestamp: When did they achieve the score.
* Score: How high was the score. 