Tomorrow:
    GameScenes are getting grabbed based on state that is in game data
    that's a map look up now instead of looking it up in a slice

Stop once all current scenes are implemented and look at refactors or designs
    aka: the prototype is done and it's time to see what we want to change
## General

add in resource manager for things like maps, sprites, objects, etc (flyweight)

factory that each scene can set I.E.: BattleSceneMega = BattleSceneMegaFactory or based on levels of the players, or the area in which they're in
    potentially use a Builder pattern

## Patterns
AI Strategy Pattern
    Patrol
    TargetHighestHealth
    TargetLowestHealth
    PreferCastingSpells
    etc...

Decorator Pattern for buffs

## Systems

Inventory in General: player has an Inventory on them that's a map[int]Item and we display that

BattleSystem -- see if we want to do Observer, or not (event bus and subscribers)

Figure out a dialogue system

Quest system would be a large map[int]Quest and GameData stores the quest index

## Scenes
Inventory Scene (CreateSceneInventory onto Game)

## Characters & Player
Character Gear: Assign slots with Items map[int]Item again?


## Renderer

do the Clear() and Present() from Engine instead of the scene?
Catch errors from Renderer interface implementations