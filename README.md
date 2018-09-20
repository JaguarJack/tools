# tools
tools website maked by vue &amp;&amp; go

## Usage
- cd tools
- npm install

## go依赖
- go get github.com/gin-gonic/gin
- go get github.com/go-sql-driver/mysql
- go get github.com/go-xorm/xorm
- go get github.com/gin-contrib/cors
- go get github.com/golang/freetype
- go get golang.org/x/image/font
- go get golang.org/x/crypto/ssh
- go get github.com/nfnt/resize
function process($start, $end)
{
	$count = 100;
	for ($i = $start; $i <= $end; $i++) {
		usleep(rand(10000, 90000));
		printf("\r installing %-100s %2d%%", str_repeat("=", $i) . ">", ($i / $count) * 100);
	}
}

function task()
{
	return 25;
}

function task2()
{
	return 50;
}

function task3()
{
	return 25;
}
static $j;
$j = task();
process(0, $j);
echo "task finished";
$end = $j + task2();
process($j, $end);
echo "task2 finished";
process($end, $end + task3());
echo "task3 finished";



