<?php

declare(strict_types=1);

namespace AOCTest\Year2022\Day02;

use AOC\PartEnum;
use AOC\Year2022\Day01\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput(): string
    {
        return "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";
    }

    /** @test */
    public function it_will_get_the_answer_for_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        //7000 + 8000 + 9000 = 24000 by the 4th group (elf)
        $this->assertSame(24000, $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(45000, $main->run(PartEnum::P2));
    }
}
