import { Recipe } from 'src/modules/recipe/entities/recipe.entity';
import {
  Column,
  Entity,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { BaseEntity } from '../../../common/entities/base.entity';

@Entity()
export class Instruction extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @ManyToOne(() => Recipe, (recipe) => recipe.instructions, {
    onDelete: 'CASCADE',
  })
  @JoinColumn({ name: 'recipe_id' })
  recipe: Recipe;

  @Column({ name: 'step_number' })
  stepNumber: number;

  @Column()
  description: string;

  @Column({ name: 'image_url' })
  imageUrl: string;
}
