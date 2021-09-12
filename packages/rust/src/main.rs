use indicatif::{ProgressBar, ProgressStyle};
use structopt::StructOpt;
///! name:          climath
///! author:        dual-lab@yandex.com
///!
///!  
///! Command line interface that has some algebriac function, developed
///! reading an cryptographic intro book.
mod cli;

fn main() {
    let _opt = cli::Opt::from_args();

    let spinner = ProgressBar::new_spinner();
    spinner.enable_steady_tick(120);
    spinner.set_style(
        ProgressStyle::default_spinner()
            .tick_strings(&[
                "ρββββββ",
                "βρβββββ",
                "ββρββββ",
                "βββρβββ",
                "ββββρββ",
                "βββββρβ",
                "ββββββρ",
            ])
            .template("{spinner} {msg}"),
    );
    spinner.finish_with_message("Done");
    spinner.disable_steady_tick();
}


