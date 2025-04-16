Changelog:
<details>
    <summary> Changelog 4/15/2025:
        -Tried adding sidebar (still bugged)
        -Tried to add favicon (failed)
        -Added working buttons so you can go between pages (!!!!)
        -Formatting improvments (proofreading, spellchecking etc. etc.)
        -Find anything that sucks? TELL ME.
    </summary>
</details>
# GOTTHStackTutorial
Hello! Welcome to my GOTTH stack tutorial. I made this as a central repo for GOTTH stack work, and for everything that's relevant.

Tech used: Go, Tailwind, HTMX, Go Templates deployed onto Google cloud (using CI/CD and cloud build/run).

Time taken: Mar 3rd - Mar 8th (maybe 40-50 hours), with most of that being spent on content.

It's got a lot of information, and touches on a ton of different topics.

TODO: (Contributors needed! Including: You! ).
1. Look over the links in the novice sections, and add more relevant ones. Any recommendations?
2. Formatting and proofreading. I made some blurbs and had chatGPT reword them into something simple and straightforward.
3. Find bugs in code (works on my machine).
4. Implement advanced sections.
5. (Difficult) Light mode revamp. We'd have to go through ALL the .html files and revamp the classes to enable light mode. Now, this is possible using an LLM to handle the find and replace - but it'll take a bit.
6. (Impossible) An "Idiomatic" or "best practices" section where we discuss the best way to do things in terms of perfomance and flexibility. This would include discussions and code examples.
7. (Tedious) Fix all the buttons in the pages to have them link to other ones.
8. (?? Easy or hard IDK) Side bar for other tutorials in a specific section.
9. (Easy, but takes a bit) Changing all parsetemplates functions to instead parseglob
10. (Hard, for me at least) Change installing js scripts from CDN to installing from a static location

-How do I run this locally?
Easy.
1. Download this project as a zip, and unzip it
2. Open the folder in terminal
3. "go mod tidy" and "go run ."
boom. there u go. Can't run? LMK!
