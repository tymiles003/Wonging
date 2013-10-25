Wonging
=======

####(WORK IN PROGRESS!)
###What is it
A little card counting tool/simulator, mainly to help myself and my friend(s) prepare our minds

Wonging: Back-counting, consists of standing behind a blackjack table that other players are playing on, and counting the cards as they are dealt. Stanford Wong first proposed the idea of back-counting, and the term "Wong" comes from his name.

###Plan
The plan is to create a complete simulator of the backjack table section of a typical casino, from the random # of players per table, to randomized counting strategies each player use or not use. The goal is to have a bird-eye view on casino operations and see how much each player would lose/win base on different strategies used. The simulator will also calculate winning chance for each hand for each player, in order to statistically determine when is the right time to start joining a game or leave a game as a player

The second phase of the project will cover group counting/strategies, by combining players into groups, this simulator hopefully may reveal different(better) winning percentages than individual operations

###Detailed plan proposal:


####Casino
---
Casino contains multiple tables, idleDealers, idlePlayers, as well as a bank to store all its cash.
The main function of casino is to assign dealers and players to different tables and control their state. It also maintains the bank, from where the owner will know how much he is profiting

####Table
---
Table contains:
*1 dealer
*multiple players
*multiple observers
*a reference to parent casino
