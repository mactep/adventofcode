DAY=$1
COOKIE="session=53616c7465645f5f5652e18885e113c2cb03d5974396e9e52318bf4a379a9a05c4c3dbb1aed430b8e659203e761c78536904f5457f6a959651c0c216fa2e9e8f"
curl --cookie $COOKIE https://adventofcode.com/2022/day/$DAY/input -o day$DAY/input
