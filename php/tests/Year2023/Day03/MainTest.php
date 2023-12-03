<?php

declare(strict_types=1);

namespace AOCTest\Year2023\Day03;

use AOC\PartEnum;
use AOC\Year2023\Day03\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput1(): string
    {
        return "
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
";
    }

    private function getTestInput2(): string
    {
        return "
12.......*..
+.........34
.......-12..
..78........
..*....60...
78.........9
.5.....23..$
8...90*12...
............
2.2......12.
.*.........*
1.1..503+.56
";
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_03_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput1());

        $this->assertSame(4361, $main->run(PartEnum::P1));

        $main->setInput($this->getTestInput2());

        $this->assertSame(925, $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_03_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput1());

        $this->assertSame(467835, $main->run(PartEnum::P2));

        $main->setInput($this->getTestInput2());

        $this->assertSame(6756, $main->run(PartEnum::P2));
    }
}
