CREATE TABLE "Cache_DefaultCard" (
  "ScryfallID" TEXT NOT NULL UNIQUE,
  "Name" TEXT NOT NULL,
  "SetName" TEXT NOT NULL,
  "CollectorNumber" TEXT NOT NULL,
  "Language" TEXT NOT NULL,
  "ReleasedAt" TEXT NOT NULL,
  "ImageSmallURL" TEXT,
  "ImageNormalURL" TEXT,
  "ManaCost" TEXT,
  "HasFoil" INTEGER NOT NULL,
  "HasNormal" INTEGER NOT NULL,
  "GathererURL" TEXT,
  "PriceNormalUSD" INTEGER,
  "PriceFoilUSD" INTEGER
);

CREATE INDEX "Index_Cache_DefaultCard_Name" ON "Cache_DefaultCard" ("Name");
CREATE INDEX "Index_Cache_DefaultCard_SetName" ON "Cache_DefaultCard" ("SetName");
