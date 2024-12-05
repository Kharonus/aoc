use std::error::Error;

pub fn solve_first_star(input: &str) -> Result<String, Box<dyn Error>> {
    let (rules, updates) = parse_input(input)?;

    let sum = updates.iter().fold(0, |acc, update| {
        if has_correct_order(update, &rules) {
            return acc + update[update.len() / 2];
        }

        acc
    });

    Ok(format!("{}", sum))
}

pub fn solve_second_star(input: &str) -> Result<String, Box<dyn Error>> {
    let (rules, updates) = parse_input(input)?;

    let sum = updates
        .iter()
        .filter_map(|update| {
            if has_correct_order(update, &rules) {
                None
            } else {
                reorder(update, &rules)
            }
        })
        .fold(0, |acc, update| acc + update[update.len() / 2]);

    Ok(format!("{}", sum))
}

#[derive(Debug)]
struct Rule(usize, usize);

impl From<Vec<usize>> for Rule {
    fn from(v: Vec<usize>) -> Self {
        Rule(v[0], v[1])
    }
}

type Update = Vec<usize>;

fn parse_input(input: &str) -> Result<(Vec<Rule>, Vec<Update>), Box<dyn Error>> {
    let mut rules_ready = false;
    let mut rules: Vec<Rule> = Vec::new();
    let mut updates: Vec<Update> = Vec::new();

    for line in input.lines() {
        if line.is_empty() {
            rules_ready = true;
            continue;
        }

        if rules_ready {
            let update: Update = line
                .split(",")
                .map(|number| number.parse())
                .collect::<Result<Update, _>>()?;

            updates.push(update);
        } else {
            let rule: Rule = line
                .split("|")
                .map(|number| number.parse())
                .take(2)
                .collect::<Result<Vec<usize>, _>>()?
                .try_into()?;

            rules.push(rule);
        }
    }

    Ok((rules, updates))
}

fn has_correct_order(update: &Update, rules: &Vec<Rule>) -> bool {
    let mut printed_pages: Vec<usize> = Vec::new();

    for page in update {
        let violation: bool = rules
            .iter()
            .filter_map(|rule| if rule.0 == *page { Some(rule.1) } else { None })
            .any(|page| printed_pages.contains(&page));

        if violation {
            return false;
        } else {
            printed_pages.push(*page);
        }
    }

    true
}

fn reorder(update: &Update, rules: &Vec<Rule>) -> Option<Update> {
    let mut ordered_update: Update = Vec::new();
    let mut applicable_rules: Vec<Rule> = filter_rules(update, &ordered_update, rules);

    while ordered_update.len() < update.len() {
        let next_page = update
            .iter()
            .filter(|page| !ordered_update.contains(page))
            .find(|page| !applicable_rules.iter().any(|rule| rule.1 == **page))?;

        ordered_update.push(*next_page);
        applicable_rules = filter_rules(update, &ordered_update, &applicable_rules);
    }

    Some(ordered_update)
}

fn filter_rules(update: &Update, ordered_update: &Update, rules: &Vec<Rule>) -> Vec<Rule> {
    rules
        .iter()
        .filter_map(|rule| {
            if ordered_update.contains(&rule.0) || ordered_update.contains(&rule.1) {
                return None;
            }

            if update.contains(&rule.0) && update.contains(&rule.1) {
                Some(Rule(rule.0, rule.1))
            } else {
                None
            }
        })
        .collect()
}
