import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Instruction } from './entities/instruction.entity';
import { InstructionController } from './instruction.controller';
import { InstructionService } from './instruction.service';

@Module({
  imports: [TypeOrmModule.forFeature([Instruction])],
  controllers: [InstructionController],
  providers: [InstructionService],
})
export class InstructionModule {}
