<?php

declare(strict_types=1);

namespace AOCTest\Year2023\Day02;

use AOC\PartEnum;
use AOC\Year2023\Day02\Main;
use PHPUnit\Framework\TestCase;

final class MainTest extends TestCase
{
    private function getTestInput(): string
    {
        return "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_02_part_1(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(8, $main->run(PartEnum::P1));
    }

    /** @test */
    public function it_will_get_the_answer_for_2023_02_part_2(): void
    {
        $main = new Main();

        $main->setInput($this->getTestInput());

        $this->assertSame(2286, $main->run(PartEnum::P2));
    }
}
