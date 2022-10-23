-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_User" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "username" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "uuid" TEXT NOT NULL,
    "skinId" TEXT,
    "capeId" TEXT,
    "currentServer" TEXT,
    "accessToken" TEXT,
    "clientToken" TEXT,
    "allowChat" BOOLEAN NOT NULL DEFAULT true,
    "allowMultiplayer" BOOLEAN NOT NULL DEFAULT true,
    "allowRealms" BOOLEAN NOT NULL DEFAULT true
);
INSERT INTO "new_User" ("accessToken", "capeId", "clientToken", "createdAt", "currentServer", "email", "id", "password", "skinId", "updatedAt", "username", "uuid") SELECT "accessToken", "capeId", "clientToken", "createdAt", "currentServer", "email", "id", "password", "skinId", "updatedAt", "username", "uuid" FROM "User";
DROP TABLE "User";
ALTER TABLE "new_User" RENAME TO "User";
CREATE UNIQUE INDEX "User_username_key" ON "User"("username");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
