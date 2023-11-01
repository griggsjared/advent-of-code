<?php

declare(strict_types=1);

namespace AOCTest\Year2022\Day03;

use AOC\PartEnum;
use AOC\Year2022\Day03\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput(): string
    {
        return "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
    }

    /** @test */
    public function it_will_get_the_answer_for_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(157, $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(70, $main->run(PartEnum::P2));
    }
}
