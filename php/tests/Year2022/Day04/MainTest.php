<?php

declare(strict_types=1);

namespace AOCTest\Year2022\Day04;

use AOC\PartEnum;
use AOC\Year2022\Day04\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput(): string
    {
        return "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";
    }

    /** @test */
    public function it_will_get_the_answer_for_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(2, $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(4, $main->run(PartEnum::P2));
    }
}
