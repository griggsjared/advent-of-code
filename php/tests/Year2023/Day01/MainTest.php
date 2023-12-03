<?php

declare(strict_types=1);

namespace AOCTest\Year2023\Day01;

use AOC\PartEnum;
use AOC\Year2023\Day01\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput1(): string
    {
        return "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";
    }

    private function getTestInput2(): string
    {
        return "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
eightwo
";
    }


    /** @test */
    public function it_will_get_the_answer_for_2023_01_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput1());

        //12 + 38 + 15 + 77 = 142
        $this->assertSame(142, $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_01_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput2());

        //29 + 83 + 13 + 24 + 42 + 14 + 76 + 82 = 363
        $this->assertSame(363, $main->run(PartEnum::P2));
    }
}
