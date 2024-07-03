// Code generated by cmd/blockhash; DO NOT EDIT.

package block

const (
	hashAir = iota
	hashAmethyst
	hashAncientDebris
	hashAndesite
	hashAnvil
	hashBanner
	hashBarrel
	hashBarrier
	hashBasalt
	hashBeacon
	hashBed
	hashBedrock
	hashBeetrootSeeds
	hashBlackstone
	hashBlastFurnace
	hashBlueIce
	hashBone
	hashBookshelf
	hashBricks
	hashCactus
	hashCake
	hashCalcite
	hashCarpet
	hashCarrot
	hashChain
	hashChest
	hashChiseledQuartz
	hashClay
	hashCoal
	hashCoalOre
	hashCobblestone
	hashCocoaBean
	hashComposter
	hashConcrete
	hashConcretePowder
	hashCopperOre
	hashCoral
	hashCoralBlock
	hashCraftingTable
	hashDeadBush
	hashDecoratedPot
	hashDeepslate
	hashDeepslateBricks
	hashDeepslateTiles
	hashDiamond
	hashDiamondOre
	hashDiorite
	hashDirt
	hashDirtPath
	hashDoubleFlower
	hashDoubleTallGrass
	hashDragonEgg
	hashDriedKelp
	hashDripstone
	hashEmerald
	hashEmeraldOre
	hashEnchantingTable
	hashEndBricks
	hashEndStone
	hashEnderChest
	hashFarmland
	hashFern
	hashFire
	hashFletchingTable
	hashFlower
	hashFroglight
	hashFurnace
	hashGlass
	hashGlassPane
	hashGlazedTerracotta
	hashGlowstone
	hashGold
	hashGoldOre
	hashGranite
	hashGrass
	hashGravel
	hashGrindstone
	hashHayBale
	hashHoneycomb
	hashInvisibleBedrock
	hashIron
	hashIronBars
	hashIronOre
	hashItemFrame
	hashJukebox
	hashKelp
	hashLadder
	hashLantern
	hashLapis
	hashLapisOre
	hashLava
	hashLeaves
	hashLectern
	hashLight
	hashLitPumpkin
	hashLog
	hashLoom
	hashMelon
	hashMelonSeeds
	hashMossCarpet
	hashMud
	hashMudBricks
	hashMuddyMangroveRoots
	hashNetherBrickFence
	hashNetherBricks
	hashNetherGoldOre
	hashNetherQuartzOre
	hashNetherSprouts
	hashNetherWart
	hashNetherWartBlock
	hashNetherite
	hashNetherrack
	hashNote
	hashObsidian
	hashPackedIce
	hashPackedMud
	hashPlanks
	hashPodzol
	hashPolishedBlackstoneBrick
	hashPotato
	hashPrismarine
	hashPumpkin
	hashPumpkinSeeds
	hashPurpur
	hashPurpurPillar
	hashQuartz
	hashQuartzBricks
	hashQuartzPillar
	hashRawCopper
	hashRawGold
	hashRawIron
	hashReinforcedDeepslate
	hashSand
	hashSandstone
	hashSeaLantern
	hashSeaPickle
	hashShortGrass
	hashShroomlight
	hashSign
	hashSkull
	hashSlab
	hashSmithingTable
	hashSmoker
	hashSnow
	hashSoulSand
	hashSoulSoil
	hashSponge
	hashSporeBlossom
	hashStainedGlass
	hashStainedGlassPane
	hashStainedTerracotta
	hashStairs
	hashStone
	hashStoneBricks
	hashStonecutter
	hashSugarCane
	hashTNT
	hashTerracotta
	hashTorch
	hashTuff
	hashWall
	hashWater
	hashWheatSeeds
	hashWood
	hashWoodDoor
	hashWoodFence
	hashWoodFenceGate
	hashWoodTrapdoor
	hashWool
	hashCustomBlockBase
)

// customBlockBase represents the base hash for all custom blocks.
var customBlockBase = uint64(hashCustomBlockBase - 1)

// NextHash returns the next free hash for custom blocks.
func NextHash() uint64 {
	customBlockBase++
	return customBlockBase
}

// Hash ...
func (Air) Hash() uint64 {
	return hashAir
}

// Hash ...
func (Amethyst) Hash() uint64 {
	return hashAmethyst
}

// Hash ...
func (AncientDebris) Hash() uint64 {
	return hashAncientDebris
}

// Hash ...
func (a Andesite) Hash() uint64 {
	return hashAndesite | uint64(boolByte(a.Polished))<<8
}

// Hash ...
func (a Anvil) Hash() uint64 {
	return hashAnvil | uint64(a.Type.Uint8())<<8 | uint64(a.Facing)<<10
}

// Hash ...
func (b Banner) Hash() uint64 {
	return hashBanner | uint64(b.Attach.Uint8())<<8
}

// Hash ...
func (b Barrel) Hash() uint64 {
	return hashBarrel | uint64(b.Facing)<<8 | uint64(boolByte(b.Open))<<11
}

// Hash ...
func (Barrier) Hash() uint64 {
	return hashBarrier
}

// Hash ...
func (b Basalt) Hash() uint64 {
	return hashBasalt | uint64(boolByte(b.Polished))<<8 | uint64(b.Axis)<<9
}

// Hash ...
func (Beacon) Hash() uint64 {
	return hashBeacon
}

// Hash ...
func (b Bed) Hash() uint64 {
	return hashBed | uint64(b.Facing)<<8 | uint64(boolByte(b.HeadSide))<<10
}

// Hash ...
func (b Bedrock) Hash() uint64 {
	return hashBedrock | uint64(boolByte(b.InfiniteBurning))<<8
}

// Hash ...
func (b BeetrootSeeds) Hash() uint64 {
	return hashBeetrootSeeds | uint64(b.Growth)<<8
}

// Hash ...
func (b Blackstone) Hash() uint64 {
	return hashBlackstone | uint64(b.Type.Uint8())<<8
}

// Hash ...
func (b BlastFurnace) Hash() uint64 {
	return hashBlastFurnace | uint64(b.Facing)<<8 | uint64(boolByte(b.Lit))<<10
}

// Hash ...
func (BlueIce) Hash() uint64 {
	return hashBlueIce
}

// Hash ...
func (b Bone) Hash() uint64 {
	return hashBone | uint64(b.Axis)<<8
}

// Hash ...
func (Bookshelf) Hash() uint64 {
	return hashBookshelf
}

// Hash ...
func (Bricks) Hash() uint64 {
	return hashBricks
}

// Hash ...
func (c Cactus) Hash() uint64 {
	return hashCactus | uint64(c.Age)<<8
}

// Hash ...
func (c Cake) Hash() uint64 {
	return hashCake | uint64(c.Bites)<<8
}

// Hash ...
func (Calcite) Hash() uint64 {
	return hashCalcite
}

// Hash ...
func (c Carpet) Hash() uint64 {
	return hashCarpet | uint64(c.Colour.Uint8())<<8
}

// Hash ...
func (c Carrot) Hash() uint64 {
	return hashCarrot | uint64(c.Growth)<<8
}

// Hash ...
func (c Chain) Hash() uint64 {
	return hashChain | uint64(c.Axis)<<8
}

// Hash ...
func (c Chest) Hash() uint64 {
	return hashChest | uint64(c.Facing)<<8
}

// Hash ...
func (ChiseledQuartz) Hash() uint64 {
	return hashChiseledQuartz
}

// Hash ...
func (Clay) Hash() uint64 {
	return hashClay
}

// Hash ...
func (Coal) Hash() uint64 {
	return hashCoal
}

// Hash ...
func (c CoalOre) Hash() uint64 {
	return hashCoalOre | uint64(c.Type.Uint8())<<8
}

// Hash ...
func (c Cobblestone) Hash() uint64 {
	return hashCobblestone | uint64(boolByte(c.Mossy))<<8
}

// Hash ...
func (c CocoaBean) Hash() uint64 {
	return hashCocoaBean | uint64(c.Facing)<<8 | uint64(c.Age)<<10
}

// Hash ...
func (c Composter) Hash() uint64 {
	return hashComposter | uint64(c.Level)<<8
}

// Hash ...
func (c Concrete) Hash() uint64 {
	return hashConcrete | uint64(c.Colour.Uint8())<<8
}

// Hash ...
func (c ConcretePowder) Hash() uint64 {
	return hashConcretePowder | uint64(c.Colour.Uint8())<<8
}

// Hash ...
func (c CopperOre) Hash() uint64 {
	return hashCopperOre | uint64(c.Type.Uint8())<<8
}

// Hash ...
func (c Coral) Hash() uint64 {
	return hashCoral | uint64(c.Type.Uint8())<<8 | uint64(boolByte(c.Dead))<<11
}

// Hash ...
func (c CoralBlock) Hash() uint64 {
	return hashCoralBlock | uint64(c.Type.Uint8())<<8 | uint64(boolByte(c.Dead))<<11
}

// Hash ...
func (CraftingTable) Hash() uint64 {
	return hashCraftingTable
}

// Hash ...
func (DeadBush) Hash() uint64 {
	return hashDeadBush
}

// Hash ...
func (p DecoratedPot) Hash() uint64 {
	return hashDecoratedPot | uint64(p.Facing)<<8
}

// Hash ...
func (d Deepslate) Hash() uint64 {
	return hashDeepslate | uint64(d.Type.Uint8())<<8 | uint64(d.Axis)<<10
}

// Hash ...
func (d DeepslateBricks) Hash() uint64 {
	return hashDeepslateBricks | uint64(boolByte(d.Cracked))<<8
}

// Hash ...
func (d DeepslateTiles) Hash() uint64 {
	return hashDeepslateTiles | uint64(boolByte(d.Cracked))<<8
}

// Hash ...
func (Diamond) Hash() uint64 {
	return hashDiamond
}

// Hash ...
func (d DiamondOre) Hash() uint64 {
	return hashDiamondOre | uint64(d.Type.Uint8())<<8
}

// Hash ...
func (d Diorite) Hash() uint64 {
	return hashDiorite | uint64(boolByte(d.Polished))<<8
}

// Hash ...
func (d Dirt) Hash() uint64 {
	return hashDirt | uint64(boolByte(d.Coarse))<<8
}

// Hash ...
func (DirtPath) Hash() uint64 {
	return hashDirtPath
}

// Hash ...
func (d DoubleFlower) Hash() uint64 {
	return hashDoubleFlower | uint64(boolByte(d.UpperPart))<<8 | uint64(d.Type.Uint8())<<9
}

// Hash ...
func (d DoubleTallGrass) Hash() uint64 {
	return hashDoubleTallGrass | uint64(boolByte(d.UpperPart))<<8 | uint64(d.Type.Uint8())<<9
}

// Hash ...
func (DragonEgg) Hash() uint64 {
	return hashDragonEgg
}

// Hash ...
func (DriedKelp) Hash() uint64 {
	return hashDriedKelp
}

// Hash ...
func (Dripstone) Hash() uint64 {
	return hashDripstone
}

// Hash ...
func (Emerald) Hash() uint64 {
	return hashEmerald
}

// Hash ...
func (e EmeraldOre) Hash() uint64 {
	return hashEmeraldOre | uint64(e.Type.Uint8())<<8
}

// Hash ...
func (EnchantingTable) Hash() uint64 {
	return hashEnchantingTable
}

// Hash ...
func (EndBricks) Hash() uint64 {
	return hashEndBricks
}

// Hash ...
func (EndStone) Hash() uint64 {
	return hashEndStone
}

// Hash ...
func (c EnderChest) Hash() uint64 {
	return hashEnderChest | uint64(c.Facing)<<8
}

// Hash ...
func (f Farmland) Hash() uint64 {
	return hashFarmland | uint64(f.Hydration)<<8
}

// Hash ...
func (Fern) Hash() uint64 {
	return hashFern
}

// Hash ...
func (f Fire) Hash() uint64 {
	return hashFire | uint64(f.Type.Uint8())<<8 | uint64(f.Age)<<9
}

// Hash ...
func (FletchingTable) Hash() uint64 {
	return hashFletchingTable
}

// Hash ...
func (f Flower) Hash() uint64 {
	return hashFlower | uint64(f.Type.Uint8())<<8
}

// Hash ...
func (f Froglight) Hash() uint64 {
	return hashFroglight | uint64(f.Type.Uint8())<<8 | uint64(f.Axis)<<10
}

// Hash ...
func (f Furnace) Hash() uint64 {
	return hashFurnace | uint64(f.Facing)<<8 | uint64(boolByte(f.Lit))<<10
}

// Hash ...
func (Glass) Hash() uint64 {
	return hashGlass
}

// Hash ...
func (GlassPane) Hash() uint64 {
	return hashGlassPane
}

// Hash ...
func (t GlazedTerracotta) Hash() uint64 {
	return hashGlazedTerracotta | uint64(t.Colour.Uint8())<<8 | uint64(t.Facing)<<12
}

// Hash ...
func (Glowstone) Hash() uint64 {
	return hashGlowstone
}

// Hash ...
func (Gold) Hash() uint64 {
	return hashGold
}

// Hash ...
func (g GoldOre) Hash() uint64 {
	return hashGoldOre | uint64(g.Type.Uint8())<<8
}

// Hash ...
func (g Granite) Hash() uint64 {
	return hashGranite | uint64(boolByte(g.Polished))<<8
}

// Hash ...
func (Grass) Hash() uint64 {
	return hashGrass
}

// Hash ...
func (Gravel) Hash() uint64 {
	return hashGravel
}

// Hash ...
func (g Grindstone) Hash() uint64 {
	return hashGrindstone | uint64(g.Attach.Uint8())<<8 | uint64(g.Facing)<<10
}

// Hash ...
func (h HayBale) Hash() uint64 {
	return hashHayBale | uint64(h.Axis)<<8
}

// Hash ...
func (Honeycomb) Hash() uint64 {
	return hashHoneycomb
}

// Hash ...
func (InvisibleBedrock) Hash() uint64 {
	return hashInvisibleBedrock
}

// Hash ...
func (Iron) Hash() uint64 {
	return hashIron
}

// Hash ...
func (IronBars) Hash() uint64 {
	return hashIronBars
}

// Hash ...
func (i IronOre) Hash() uint64 {
	return hashIronOre | uint64(i.Type.Uint8())<<8
}

// Hash ...
func (i ItemFrame) Hash() uint64 {
	return hashItemFrame | uint64(i.Facing)<<8 | uint64(boolByte(i.Glowing))<<11
}

// Hash ...
func (Jukebox) Hash() uint64 {
	return hashJukebox
}

// Hash ...
func (k Kelp) Hash() uint64 {
	return hashKelp | uint64(k.Age)<<8
}

// Hash ...
func (l Ladder) Hash() uint64 {
	return hashLadder | uint64(l.Facing)<<8
}

// Hash ...
func (l Lantern) Hash() uint64 {
	return hashLantern | uint64(boolByte(l.Hanging))<<8 | uint64(l.Type.Uint8())<<9
}

// Hash ...
func (Lapis) Hash() uint64 {
	return hashLapis
}

// Hash ...
func (l LapisOre) Hash() uint64 {
	return hashLapisOre | uint64(l.Type.Uint8())<<8
}

// Hash ...
func (l Lava) Hash() uint64 {
	return hashLava | uint64(boolByte(l.Still))<<8 | uint64(l.Depth)<<9 | uint64(boolByte(l.Falling))<<17
}

// Hash ...
func (l Leaves) Hash() uint64 {
	return hashLeaves | uint64(l.Wood.Uint8())<<8 | uint64(boolByte(l.Persistent))<<12 | uint64(boolByte(l.ShouldUpdate))<<13
}

// Hash ...
func (l Lectern) Hash() uint64 {
	return hashLectern | uint64(l.Facing)<<8
}

// Hash ...
func (l Light) Hash() uint64 {
	return hashLight | uint64(l.Level)<<8
}

// Hash ...
func (l LitPumpkin) Hash() uint64 {
	return hashLitPumpkin | uint64(l.Facing)<<8
}

// Hash ...
func (l Log) Hash() uint64 {
	return hashLog | uint64(l.Wood.Uint8())<<8 | uint64(boolByte(l.Stripped))<<12 | uint64(l.Axis)<<13
}

// Hash ...
func (l Loom) Hash() uint64 {
	return hashLoom | uint64(l.Facing)<<8
}

// Hash ...
func (Melon) Hash() uint64 {
	return hashMelon
}

// Hash ...
func (m MelonSeeds) Hash() uint64 {
	return hashMelonSeeds | uint64(m.Growth)<<8 | uint64(m.Direction)<<16
}

// Hash ...
func (MossCarpet) Hash() uint64 {
	return hashMossCarpet
}

// Hash ...
func (Mud) Hash() uint64 {
	return hashMud
}

// Hash ...
func (MudBricks) Hash() uint64 {
	return hashMudBricks
}

// Hash ...
func (m MuddyMangroveRoots) Hash() uint64 {
	return hashMuddyMangroveRoots | uint64(m.Axis)<<8
}

// Hash ...
func (NetherBrickFence) Hash() uint64 {
	return hashNetherBrickFence
}

// Hash ...
func (n NetherBricks) Hash() uint64 {
	return hashNetherBricks | uint64(n.Type.Uint8())<<8
}

// Hash ...
func (NetherGoldOre) Hash() uint64 {
	return hashNetherGoldOre
}

// Hash ...
func (NetherQuartzOre) Hash() uint64 {
	return hashNetherQuartzOre
}

// Hash ...
func (NetherSprouts) Hash() uint64 {
	return hashNetherSprouts
}

// Hash ...
func (n NetherWart) Hash() uint64 {
	return hashNetherWart | uint64(n.Age)<<8
}

// Hash ...
func (n NetherWartBlock) Hash() uint64 {
	return hashNetherWartBlock | uint64(boolByte(n.Warped))<<8
}

// Hash ...
func (Netherite) Hash() uint64 {
	return hashNetherite
}

// Hash ...
func (Netherrack) Hash() uint64 {
	return hashNetherrack
}

// Hash ...
func (Note) Hash() uint64 {
	return hashNote
}

// Hash ...
func (o Obsidian) Hash() uint64 {
	return hashObsidian | uint64(boolByte(o.Crying))<<8
}

// Hash ...
func (PackedIce) Hash() uint64 {
	return hashPackedIce
}

// Hash ...
func (PackedMud) Hash() uint64 {
	return hashPackedMud
}

// Hash ...
func (p Planks) Hash() uint64 {
	return hashPlanks | uint64(p.Wood.Uint8())<<8
}

// Hash ...
func (Podzol) Hash() uint64 {
	return hashPodzol
}

// Hash ...
func (b PolishedBlackstoneBrick) Hash() uint64 {
	return hashPolishedBlackstoneBrick | uint64(boolByte(b.Cracked))<<8
}

// Hash ...
func (p Potato) Hash() uint64 {
	return hashPotato | uint64(p.Growth)<<8
}

// Hash ...
func (p Prismarine) Hash() uint64 {
	return hashPrismarine | uint64(p.Type.Uint8())<<8
}

// Hash ...
func (p Pumpkin) Hash() uint64 {
	return hashPumpkin | uint64(boolByte(p.Carved))<<8 | uint64(p.Facing)<<9
}

// Hash ...
func (p PumpkinSeeds) Hash() uint64 {
	return hashPumpkinSeeds | uint64(p.Growth)<<8 | uint64(p.Direction)<<16
}

// Hash ...
func (Purpur) Hash() uint64 {
	return hashPurpur
}

// Hash ...
func (p PurpurPillar) Hash() uint64 {
	return hashPurpurPillar | uint64(p.Axis)<<8
}

// Hash ...
func (q Quartz) Hash() uint64 {
	return hashQuartz | uint64(boolByte(q.Smooth))<<8
}

// Hash ...
func (QuartzBricks) Hash() uint64 {
	return hashQuartzBricks
}

// Hash ...
func (q QuartzPillar) Hash() uint64 {
	return hashQuartzPillar | uint64(q.Axis)<<8
}

// Hash ...
func (RawCopper) Hash() uint64 {
	return hashRawCopper
}

// Hash ...
func (RawGold) Hash() uint64 {
	return hashRawGold
}

// Hash ...
func (RawIron) Hash() uint64 {
	return hashRawIron
}

// Hash ...
func (ReinforcedDeepslate) Hash() uint64 {
	return hashReinforcedDeepslate
}

// Hash ...
func (s Sand) Hash() uint64 {
	return hashSand | uint64(boolByte(s.Red))<<8
}

// Hash ...
func (s Sandstone) Hash() uint64 {
	return hashSandstone | uint64(s.Type.Uint8())<<8 | uint64(boolByte(s.Red))<<10
}

// Hash ...
func (SeaLantern) Hash() uint64 {
	return hashSeaLantern
}

// Hash ...
func (s SeaPickle) Hash() uint64 {
	return hashSeaPickle | uint64(s.AdditionalCount)<<8 | uint64(boolByte(s.Dead))<<16
}

// Hash ...
func (ShortGrass) Hash() uint64 {
	return hashShortGrass
}

// Hash ...
func (Shroomlight) Hash() uint64 {
	return hashShroomlight
}

// Hash ...
func (s Sign) Hash() uint64 {
	return hashSign | uint64(s.Wood.Uint8())<<8 | uint64(s.Attach.Uint8())<<12
}

// Hash ...
func (s Skull) Hash() uint64 {
	return hashSkull | uint64(s.Attach.FaceUint8())<<8
}

// Hash ...
func (s Slab) Hash() uint64 {
	return hashSlab | s.Block.Hash()<<8 | uint64(boolByte(s.Top))<<24 | uint64(boolByte(s.Double))<<25
}

// Hash ...
func (SmithingTable) Hash() uint64 {
	return hashSmithingTable
}

// Hash ...
func (s Smoker) Hash() uint64 {
	return hashSmoker | uint64(s.Facing)<<8 | uint64(boolByte(s.Lit))<<10
}

// Hash ...
func (Snow) Hash() uint64 {
	return hashSnow
}

// Hash ...
func (SoulSand) Hash() uint64 {
	return hashSoulSand
}

// Hash ...
func (SoulSoil) Hash() uint64 {
	return hashSoulSoil
}

// Hash ...
func (s Sponge) Hash() uint64 {
	return hashSponge | uint64(boolByte(s.Wet))<<8
}

// Hash ...
func (SporeBlossom) Hash() uint64 {
	return hashSporeBlossom
}

// Hash ...
func (g StainedGlass) Hash() uint64 {
	return hashStainedGlass | uint64(g.Colour.Uint8())<<8
}

// Hash ...
func (p StainedGlassPane) Hash() uint64 {
	return hashStainedGlassPane | uint64(p.Colour.Uint8())<<8
}

// Hash ...
func (t StainedTerracotta) Hash() uint64 {
	return hashStainedTerracotta | uint64(t.Colour.Uint8())<<8
}

// Hash ...
func (s Stairs) Hash() uint64 {
	return hashStairs | s.Block.Hash()<<8 | uint64(boolByte(s.UpsideDown))<<24 | uint64(s.Facing)<<25
}

// Hash ...
func (s Stone) Hash() uint64 {
	return hashStone | uint64(boolByte(s.Smooth))<<8
}

// Hash ...
func (s StoneBricks) Hash() uint64 {
	return hashStoneBricks | uint64(s.Type.Uint8())<<8
}

// Hash ...
func (s Stonecutter) Hash() uint64 {
	return hashStonecutter | uint64(s.Facing)<<8
}

// Hash ...
func (c SugarCane) Hash() uint64 {
	return hashSugarCane | uint64(c.Age)<<8
}

// Hash ...
func (TNT) Hash() uint64 {
	return hashTNT
}

// Hash ...
func (Terracotta) Hash() uint64 {
	return hashTerracotta
}

// Hash ...
func (t Torch) Hash() uint64 {
	return hashTorch | uint64(t.Facing)<<8 | uint64(t.Type.Uint8())<<11
}

// Hash ...
func (Tuff) Hash() uint64 {
	return hashTuff
}

// Hash ...
func (w Wall) Hash() uint64 {
	return hashWall | w.Block.Hash()<<8 | uint64(w.NorthConnection.Uint8())<<24 | uint64(w.EastConnection.Uint8())<<26 | uint64(w.SouthConnection.Uint8())<<28 | uint64(w.WestConnection.Uint8())<<30 | uint64(boolByte(w.Post))<<32
}

// Hash ...
func (w Water) Hash() uint64 {
	return hashWater | uint64(boolByte(w.Still))<<8 | uint64(w.Depth)<<9 | uint64(boolByte(w.Falling))<<17
}

// Hash ...
func (s WheatSeeds) Hash() uint64 {
	return hashWheatSeeds | uint64(s.Growth)<<8
}

// Hash ...
func (w Wood) Hash() uint64 {
	return hashWood | uint64(w.Wood.Uint8())<<8 | uint64(boolByte(w.Stripped))<<12 | uint64(w.Axis)<<13
}

// Hash ...
func (d WoodDoor) Hash() uint64 {
	return hashWoodDoor | uint64(d.Wood.Uint8())<<8 | uint64(d.Facing)<<12 | uint64(boolByte(d.Open))<<14 | uint64(boolByte(d.Top))<<15 | uint64(boolByte(d.Right))<<16
}

// Hash ...
func (w WoodFence) Hash() uint64 {
	return hashWoodFence | uint64(w.Wood.Uint8())<<8
}

// Hash ...
func (f WoodFenceGate) Hash() uint64 {
	return hashWoodFenceGate | uint64(f.Wood.Uint8())<<8 | uint64(f.Facing)<<12 | uint64(boolByte(f.Open))<<14 | uint64(boolByte(f.Lowered))<<15
}

// Hash ...
func (t WoodTrapdoor) Hash() uint64 {
	return hashWoodTrapdoor | uint64(t.Wood.Uint8())<<8 | uint64(t.Facing)<<12 | uint64(boolByte(t.Open))<<14 | uint64(boolByte(t.Top))<<15
}

// Hash ...
func (w Wool) Hash() uint64 {
	return hashWool | uint64(w.Colour.Uint8())<<8
}
