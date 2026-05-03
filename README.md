# BugCraft

## Pitch

BugCraft is a cozy terminal break game where developers grow resources while working, craft spells during short breaks, and raid haunted codebases to hunt bugs before the break timer closes the dungeon.

```txt
Cronfarm → Stacksmith → Regex Rogue → Bug Hunt → Back to Work
```

---

## Core Idea

BugCraft combines four developer-friendly mini-systems into one 10–15 minute command-line game loop:

1. **Cronfarm** — passive resource gathering while the player is away.
2. **Stacksmith** — stack-based crafting for spells, tools, and consumables.
3. **Regex Rogue** — small regex puzzles that tune or empower spells.
4. **Bug Hunt** — short roguelike raids against software bugs inside cursed codebases.

The key design constraint is intentional: **BugCraft should be fun, but not endlessly addictive.** It is built around healthy work breaks. When the break timer ends, the game locks the run and nudges the player back to work.

### Designed for 10–15 Minute Breaks

BugCraft should feel satisfying in a short session.

A player should be able to:

- Harvest resources.
- Craft one or two useful spells.
- Solve one small regex puzzle.
- Complete or attempt one bughunt raid.
- See a clear session summary.


---

## Game Modes

# 1. Cronfarm

Cronfarm is the passive resource gathering layer.

The player plants developer-themed crops. These mature over real elapsed time while the player is working.

Example crops:

| Crop | Resource Produced | Use |
|---|---|---|
| Coffee Bean | Coffee | Restore HP or focus |
| Regex Root | Regex Essence | Power regex spells |
| Stack Crystal | Stack Ore | Craft stack spells |
| Heap Mushroom | Heap Spores | Memory-related spells |
| Mutex Mint | Mutex Leaf | Counter RaceCondition enemies |

Example farm screen:

```txt
┌────────────── BUGCRAFT FARM ───────────────┐
│ [c] Coffee Beans      ready now            │
│ [r] Regex Roots       ready in 04:12       │
│ [s] Stack Crystal     ready in 11:05       │
│ [h] Heap Mushroom     ready now            │
│ [m] Mutex Mint        empty plot           │
└────────────────────────────────────────────┘
```

Example commands:

```bash
bugcraft farm
bugcraft plant coffee
bugcraft harvest
```

Behavior:

- Crops have a `planted_at` timestamp.
- Crops mature after `growth_minutes`.
- Harvest adds resources to inventory.
- The project can simulate cron behavior using timestamps instead of system cron.

---

# 2. Stacksmith

Stacksmith is the crafting system.

Players push resources onto a stack and combine them into spells or items.

Example:

```txt
Stack:
[ Coffee Bean  ]
[ Regex Root   ]
[ Stack Crystal]

> combine
```

Result:

```txt
Crafted: Caffeinated Pattern Blast
Effect: Deal bonus damage to enemies matching a regex pattern.
```

Example commands:

```bash
bugcraft craft
bugcraft push regex-root
bugcraft push coffee
bugcraft combine
```

Example recipes:

| Recipe | Ingredients | Effect |
|---|---|---|
| Caffeinated Pattern Blast | Coffee + Regex Root + Stack Crystal | Regex-powered attack spell |
| Mutex Trap | Mutex Mint + Stack Crystal | Prevents RaceCondition from double-attacking |
| Heap Lantern | Heap Mushroom + Coffee | Reveals hidden MemoryLeak enemies |
| Null Shield | Stack Crystal + Coffee | Blocks one NullPointer attack |
| Grep Charm | Regex Root + Mutex Mint | Improves regex puzzle reward |

v0.1 simplification:

- Stack crafting can be implemented as an ordered list.
- Recipes can be exact ordered matches.
- Later versions can support flexible or partial recipes.

---

# 3. Regex Rogue

Regex Rogue is the puzzle layer.

Players solve small pattern-matching puzzles to empower spells before entering a bug dungeon.

Example puzzle:

```txt
Target only error logs:

[ERR] cache overflow
[INFO] cron completed
[ERR] null pointer
[WARN] slow query

Regex:
> ^\[ERR\].*
```

Result:

```txt
Success!
Caffeinated Pattern Blast gains +2 damage for this raid.
```

Example puzzle types:

- Match only error lines
- Extract bug names
- Ignore info/debug logs
- Capture test names
- Match valid stack traces

Example scoring:

- Perfect match: full bonus
- Partial match: small bonus
- Overmatch: no bonus or minor penalty

Community expansion:

- Regex puzzle packs
- Language-specific log formats
- Regex golf challenges
- Evil input strings
- Timed puzzles

---

# 4. Bug Hunt

Bug Hunt is the short roguelike combat mode.

The player enters a haunted codebase and fights bugs on a small ASCII grid.

Example dungeon:

```txt
┌──────────────────────────┐
│ @ . . B . . C . .        │
│ . # . . . # . . .        │
│ . . F . . . B . .        │
│ . . . . # . . . M        │
└──────────────────────────┘

@ You
B Bug
C Coffee
F Failing Test
M MemoryLeak
# Wall
```

Controls:

```txt
w/a/s/d     move
i           inspect
f           fix/attack
c           cast spell
q           quit raid
```

Example enemies:

| Enemy | Behavior |
|---|---|
| NullPointer | Surprise burst damage |
| RaceCondition | May attack twice |
| MemoryLeak | Gets stronger every few turns |
| OffByOne | Dodges first attack |
| RegexBacktrackingDemon | Punishes failed regex spell |
| LegacyMonolith | Boss |

Example boss text:

```txt
LegacyMonolith appears.
It has 10,000 lines of undocumented behavior.
It remembers every shortcut you ever shipped.
```

Win condition:

- Defeat the boss or fix a target number of bugs before the break timer ends.

Fail condition:

- HP reaches zero.
- Break timer expires.
- Player exits raid.

---