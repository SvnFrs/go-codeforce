#include <stdio.h>

int isdigit();
int main(void)
{
    char direction[1000];
    scanf("%s", direction);
    int x = 0, y = 0;
    for (int i = 0; direction[i] != '\0'; i++)
    {
        if (direction[i] == 'N')
        {
            y++;
        }
        else if (isdigit(direction[i])) {
            int num = direction[i] - '0';
            if (direction[i+1] == 'N')
            {
                y += num;
            }
            else if (direction[i+1] == 'S')
            {
                y -= num;
            }
            else if (direction[i+1] == 'E')
            {
                x += num;
            }
            else if (direction[i+1] == 'W')
            {
                x -= num;
            }
            i++;
        }
        else if (direction[i] == 'S')
        {
            y--;
        }
        else if (direction[i] == 'E')
        {
            x++;
        }
        else if (direction[i] == 'W')
        {
            x--;
        }
    }
    printf("%d %d", x, y);
}
