<?php

declare(strict_types=1);

namespace AOCTest\Year2023\Day06;

use AOC\PartEnum;
use AOC\Year2023\Day06\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput(): string
    {
        return "
Time:      7  15   30
Distance:  9  40  200
";
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_06_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(288, $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_06_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(71503, $main->run(PartEnum::P2));
    }
}
