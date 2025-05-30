import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateInstructionDto } from './dto/create-instruction.dto';
import { UpdateInstructionDto } from './dto/update-instruction.dto';
import { Instruction } from './entities/instruction.entity';

@Injectable()
export class InstructionService {
  constructor(
    @InjectRepository(Instruction)
    private instructionRepository: Repository<Instruction>,
  ) {}

  create(createInstructionDto: CreateInstructionDto) {
    return 'This action adds a new instruction';
  }

  findAll() {
    return this.instructionRepository.find();
  }

  findOne(id: number) {
    return `This action returns a #${id} instruction`;
  }

  update(id: number, updateInstructionDto: UpdateInstructionDto) {
    return `This action updates a #${id} instruction`;
  }

  remove(id: number) {
    return `This action removes a #${id} instruction`;
  }
}
