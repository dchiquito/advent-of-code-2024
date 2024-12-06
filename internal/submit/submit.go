package submit

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

/*
<article><p>That's not the right answer; your answer is too low.  If you're stuck, make sure you're using the full input data; there are also some general tips on the <a href="/2019/about">about page</a>, or you can ask for hints on the <a href="https://www.reddit.com/r/adventofcode/" target="_blank">subreddit</a>.  Please wait one minute before trying again. <a href="/2019/day/1">[Return to Day 1]</a></p></article>
*/
/*
<article><p>That's not the right answer.  If you're stuck, make sure you're using the full input data; there are also some general tips on the <a href="/2019/about">about page</a>, or you can ask for hints on the <a href="https://www.reddit.com/r/adventofcode/" target="_blank">subreddit</a>.  Because you have guessed incorrectly 4 times on this puzzle, please wait 5 minutes before trying again. <a href="/2019/day/1">[Return to Day 1]</a></p></article>
*/
/*
<article><p>You gave an answer too recently; you have to wait after submitting an answer before trying again.  You have 31s left to wait. <a href="/2019/day/1">[Return to Day 1]</a></p></article>
*/
/*
<article><p>That's the right answer!  You are <span class="day-success">one gold star</span> closer to rescuing Santa. <a href="/2019/day/1#part2">[Continue to Part Two]</a></p></article>
*/
/*
<article><p>That's the right answer!  You are <span class="day-success">one gold star</span> closer to rescuing Santa.</p><p>You have completed Day 1! You can <span class="share">[Share<span class="share-content">on
  <a href="https://twitter.com/intent/tweet?text=I+just+completed+%22The+Tyranny+of+the+Rocket+Equation%22+%2D+Day+1+%2D+Advent+of+Code+2019&amp;url=https%3A%2F%2Fadventofcode%2Ecom%2F2019%2Fday%2F1&amp;related=ericwastl&amp;hashtags=AdventOfCode" target="_blank">Twitter</a>
  <a href="javascript:void(0);" onclick="var ms; try{ms=localStorage.getItem('mastodon.server')}finally{} if(typeof ms!=='string')ms=''; ms=prompt('Mastodon Server?',ms); if(typeof ms==='string' && ms.length){this.href='https://'+ms+'/share?text=I+just+completed+%22The+Tyranny+of+the+Rocket+Equation%22+%2D+Day+1+%2D+Advent+of+Code+2019+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2019%2Fday%2F1';try{localStorage.setItem('mastodon.server',ms);}finally{}}else{return false;}" target="_blank">Mastodon</a
></span>]</span> this victory or <a href="/2019">[Return to Your Advent Calendar]</a>.</p></article>
*/

func FindPayload(html []byte) string {
	re := regexp.MustCompile("<main>([\\w\\W]*)</main>")
	// TODO filter out the footer
	return string(re.FindSubmatch(html)[1])
}

func SendAnswer(day int, level int, answer string) {
	_url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/answer", day)
	form := url.Values{}
	form.Set("level", strconv.Itoa(level))
	form.Set("answer", answer)
	req, err := http.NewRequest("POST", _url, strings.NewReader(form.Encode()))
	util.Check(err, "Failed to set up request")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp := util.SendRequest(req)
	defer resp.Body.Close()
	html, err := io.ReadAll(resp.Body)
	util.Check(err, "Failed to read response when fetching endpoint")

	result := FindPayload(html)

	fmt.Println(result)
}
