<?php

declare(strict_types=1);

namespace AOC;

interface MainInterface
{
    public function run(PartEnum $part): string|int;

    public function setInput(string $string): void;

    public function getInput(): string;
}
