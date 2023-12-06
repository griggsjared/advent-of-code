<?php

declare(strict_types=1);

namespace AOCTest\Year2023\Day04;

use AOC\PartEnum;
use AOC\Year2023\Day04\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput(): string
    {
        return "
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
";
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_04_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(13, $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_04_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(30, $main->run(PartEnum::P2));
    }
}
