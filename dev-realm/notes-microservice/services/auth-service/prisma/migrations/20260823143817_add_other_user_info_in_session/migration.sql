/*
  Warnings:

  - You are about to drop the column `ipAddress` on the `sessions` table. All the data in the column will be lost.
  - You are about to drop the column `isDeleted` on the `sessions` table. All the data in the column will be lost.
  - You are about to drop the column `isRevoked` on the `sessions` table. All the data in the column will be lost.
  - You are about to drop the column `revokedAt` on the `sessions` table. All the data in the column will be lost.
  - You are about to drop the column `userAgent` on the `sessions` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "sessions" DROP COLUMN "ipAddress",
DROP COLUMN "isDeleted",
DROP COLUMN "isRevoked",
DROP COLUMN "revokedAt",
DROP COLUMN "userAgent",
ADD COLUMN     "ip_address" TEXT,
ADD COLUMN     "is_deleted" BOOLEAN NOT NULL DEFAULT false,
ADD COLUMN     "is_revoked" BOOLEAN NOT NULL DEFAULT false,
ADD COLUMN     "revoked_at" TIMESTAMP(3),
ADD COLUMN     "user_agent" TEXT;
