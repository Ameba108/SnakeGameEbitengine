# Змейка с помощью ebitengine
## Главное меню:

![Снимок экрана от 2024-07-08 17-51-10](https://github.com/Ameba108/SnakeGameEbitengine/assets/136710964/bd2c7bea-99e7-4568-b9f3-d989a1dcb0e3)

## Сама игра:

![Снимок экрана от 2024-07-08 17-57-30](https://github.com/Ameba108/SnakeGameEbitengine/assets/136710964/b8a866fe-4e65-4171-9eb8-3016c7c8bc61)

# Игра
Эта копия игры Змейка, только сделанная на Golang с помощью Ebitengine.
Нажав ЛКМ на "Старт" в начальном меню, игра запускается. Правила ни чем не отличаются от ороигинальной игры. Змейка управляется с помощью стрелок вверх, вниз, вправо и влево. Когда змейка съедает плоды, ее тело растет, количество очков и сокрость увеличивается. Если змейка врезается в свой хвост или в границы окна, игра заканчивается.

# Код
*Collision*

В пакете Collision находится константы размера игрового экрана, а также струтктура Vector, для передвижения объектов по оси X или Y.

*MainMenu*

В пакете MainMenu создается главное меню игры, с кнопкой и изображением.

*Game*

В пакете Game прописана логика змейки, ее передвижение, увеличение ее тела, а также создание еды. Отрисовывается задний фон.

*Main*

В пакете Main отрисовываются объекты. Также прописана функция Restart, позволяющая начать игру сначала. В Main также прописано нажатие клавиш при движении змейки, увеличение скорости змейки, а также логика поражения и конца игры.

# Установка
1) `git clone https://github.com/Ameba108/SnakeGameEbitengine`
2) `cd SnakeGameEbitengine`
3) `go run main/main.go`
   
