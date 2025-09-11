/*
  Warnings:

  - You are about to drop the column `user_id` on the `Shifts` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[name]` on the table `Users` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `user_name` to the `Shifts` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE `Shifts` DROP FOREIGN KEY `Shifts_user_id_fkey`;

-- DropIndex
DROP INDEX `Shift_user_id_fkey` ON `Shifts`;

-- AlterTable
ALTER TABLE `Shifts` DROP COLUMN `user_id`,
    ADD COLUMN `user_name` VARCHAR(191) NOT NULL;

-- CreateIndex
CREATE INDEX `Shift_user_id_fkey` ON `Shifts`(`user_name`);

-- CreateIndex
CREATE UNIQUE INDEX `Users_name_key` ON `Users`(`name`);

-- AddForeignKey
ALTER TABLE `Shifts` ADD CONSTRAINT `Shifts_user_name_fkey` FOREIGN KEY (`user_name`) REFERENCES `Users`(`name`) ON DELETE RESTRICT ON UPDATE CASCADE;
