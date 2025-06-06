import { Controller, Get, Post, Body, Patch, Param, Delete } from '@nestjs/common';
import { InstructionService } from './instruction.service';
import { CreateInstructionDto } from './dto/create-instruction.dto';
import { UpdateInstructionDto } from './dto/update-instruction.dto';

@Controller('instruction')
export class InstructionController {
  constructor(private readonly instructionService: InstructionService) {}

  @Post()
  create(@Body() createInstructionDto: CreateInstructionDto) {
    return this.instructionService.create(createInstructionDto);
  }

  @Get()
  findAll() {
    return this.instructionService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.instructionService.findOne(+id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateInstructionDto: UpdateInstructionDto) {
    return this.instructionService.update(+id, updateInstructionDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.instructionService.remove(+id);
  }
}
