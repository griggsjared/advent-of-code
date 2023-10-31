<?php

declare(strict_types=1);

namespace AOCTest\Year2022\Day01;

use AOC\PartEnum;
use AOC\Year2022\Day01\Main;
use PHPUnit\Framework\TestCase;

final class Day01Test extends TestCase
{
    /** @test */
    public function it_will_get_the_answer_for_part_1(): void
    {
        $main = new Main();

        $main->setInput('Hello World');

        $this->assertSame('Hello World', $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_part_2(): void
    {
        $main = new Main();

        $main->setInput('Hello World');

        $this->assertSame('Hello World', $main->run(PartEnum::P2));
    }
}
