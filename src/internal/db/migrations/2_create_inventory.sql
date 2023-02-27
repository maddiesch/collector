CREATE TABLE "Inventory" (
  "GroupName" TEXT NOT NULL,
  "Name" TEXT NOT NULL,
  "SetName" TEXT NOT NULL,
  "CollectorNumber" TEXT NOT NULL,
  "IsFoil" INTEGER NOT NULL,
  "Condition" TEXT NOT NULL,
  "Language" TEXT NOT NULL,
  "CreatedAt" INTEGER NOT NULL
);

CREATE INDEX "Index_Inventory_GroupName" ON "Inventory" ("GroupName");
CREATE INDEX "Index_Inventory_Name" ON "Inventory" ("Name");
CREATE INDEX "Index_Inventory_SetName" ON "Inventory" ("SetName");
CREATE INDEX "Index_Inventory_IsFoil" ON "Inventory" ("IsFoil");
CREATE INDEX "Index_Inventory_Condition" ON "Inventory" ("Condition");
