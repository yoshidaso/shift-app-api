/*
  Warnings:

  - You are about to drop the column `endAt` on the `Shifts` table. All the data in the column will be lost.
  - You are about to drop the column `startAt` on the `Shifts` table. All the data in the column will be lost.
  - You are about to drop the column `workContent` on the `Shifts` table. All the data in the column will be lost.
  - You are about to drop the column `createdAt` on the `Users` table. All the data in the column will be lost.
  - Added the required column `end_at` to the `Shifts` table without a default value. This is not possible if the table is not empty.
  - Added the required column `start_at` to the `Shifts` table without a default value. This is not possible if the table is not empty.
  - Added the required column `work_content` to the `Shifts` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE `Shifts` DROP COLUMN `endAt`,
    DROP COLUMN `startAt`,
    DROP COLUMN `workContent`,
    ADD COLUMN `end_at` DATETIME(3) NOT NULL,
    ADD COLUMN `start_at` DATETIME(3) NOT NULL,
    ADD COLUMN `work_content` VARCHAR(191) NOT NULL;

-- AlterTable
ALTER TABLE `Users` DROP COLUMN `createdAt`,
    ADD COLUMN `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3);
