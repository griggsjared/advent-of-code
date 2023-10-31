<?php

declare(strict_types=1);

namespace AOC\Year2022\Day01;

use AOC\MainInterface;
use AOC\PartEnum;
use InvalidArgumentException;

class Main implements MainInterface
{
    private string $input;

    public function __construct()
    {
        if (file_exists(__DIR__ . '/input.txt')) {
            $this->setInput(file_get_contents(__DIR__ . '/input.txt'));
        }
    }

    public function setInput(string $string): void
    {
        $this->input = $string;
    }

    public function getInput(): string
    {
        if(!isset($this->input)) {
            throw new InvalidArgumentException(
                'Input must be set either by an existing input.txt file or by setting via setInput.'
            );
        }

        return $this->input;
    }

    public function run(PartEnum $part): string
    {
        return match($part) {
            PartEnum::P1 => $this->part1(),
            PartEnum::P2 => $this->part2(),
            default => throw new InvalidArgumentException('Invalid part.'),
        };
    }

    private function part1(): string
    {
        return $this->getInput();
    }

    private function part2(): string
    {
        return $this->getInput();
    }
}
