<?php

declare(strict_types=1);

namespace AOCTest\Year2022\Day02;

use AOC\PartEnum;
use AOC\Year2022\Day02\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput(): string
    {
        return "A Y
B X
C Z";
    }

    /** @test */
    public function it_will_get_the_answer_for_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(15, $main->run(PartEnum::P1));
    }

    // /** @test */
    public function it_will_get_the_answer_for_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(12, $main->run(PartEnum::P2));
    }
}
