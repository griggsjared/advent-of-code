<?php

declare(strict_types=1);

namespace AOCTest\Year2022\Day01;

use AOC\Year2022\Day01\Main;
use PHPUnit\Framework\TestCase;

final class Day01Test extends TestCase
{
    /** @test */
    public function it_will_run_with_the_answer(): void
    {
        $dayMain = new Main();

        $this->assertSame('Hello World', $dayMain->run());
    }
}
