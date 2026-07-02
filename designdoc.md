# DesignDoc

## UI/UX
- Game size 640x360

- Grid tile is 80x40 pixel, and the battle field is 8x4 Grid. Refer to mockup.xcf

- The sprite of any size but typically 80x100 for humanoid character and is anchored on bottom center

- The aesthethic will be some sort of mild-cyberpunk with focus on kurowear aesthethic

- the portrait will have black outline refer to face1.xcf and face2.xcf

- the sprite will not have outline, or use as few as possible

## Audio

Ensure to have ogg as default audio format

Use these presets in LMMS

- Okt-String 2
- SuperSaw Lead
- HugeGrittyBass
- PowerStrings
- whatever in the cyberwave-presets
  - LEAD tube (upcoming big bad)
  - LEAD digigrid 
  - KEYS cybernetic
  - BASS Rolling
  - LEAD Neo
  - BASS Cyberline
  - BASS Neon Reese
  - PAD Glister
  - LEAD synthpunk
  - BASS Rezzolution
  - KEYS infected (curious)
  - PLUCK Journey
  - BASS Digigrid
  - PLUCK Night star
  - PLUCK Hammer (unease)
  - PAD Somewhere
  - BASS Smoth Cyber
  - Pluck Night
  - Pluck Moon
  - KEYS Retro
  - LEAD Chaos
  - BASS Breakthrough
  - BASS Tube
  - LEAD Light
  - KEYS Retro Bell
  - SQ Psy bass
  - SQ Frog Bass
  - SQ Disco Bass
  - SQ Background galaxy
  - PL Sweet
  - PL steel strings
  - PL muted guitars 2
  - LD Square lead 1
  - LD Flow
  - BA cyberpunk 1/2
  - 

## Battle system
- Let cooldown have its own counter and not just start/end
  this to allows pause

- move player character and not just jumping around
  
- have X,Y,Z component for element. Make sure the Z is the depth.

- replace gridpos with something more granular

- add proper hit detection

- add meter that allows for powerup

- The powerup will have cost and paid using orb

- Player get orb each time the meter is full or a wave is cleared

- the menus for now are: 
  * GrimoireGun
    * DamagePlus  2x
    * RateUp      2x
  * Suit
    * Heal        2x
    * Infiltrate  1x
  * GigaAtk
    * DamageAll   4x



## Exploration
