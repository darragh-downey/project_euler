import pprint
import json
    

def gen_years(start_range, end_range):
    years = []
    years.extend(range(start_range-1, end_range))
    years.append(end_range)

    months = gen_months()
    yrs = []
    leaps = [is_leap_year(y) for y in years]
    for y, l in zip(years, leaps):
        # months should be list of dict of [{name: n, days: d}...]
        # where days has been adjusted according to leap year
        days = 0
        day_sum = 0
        ms = []
        for m in months:
            if m["leap"] == True and l == True:
                days = m["days"] + 1
            else:
                days = m["days"]
            ms.append({"name": m["name"], "days": days})
            day_sum += days
        
        yrs.append({"year": y, "leap": l, "months": ms, "days": day_sum})


    return yrs


def gen_months():
    with open("./months.json") as f:
        j = json.load(f)
        return j


def gen_weekdays():
    with open("./days.json") as f:
        j = json.load(f)
        return j


def is_leap_year(year):
    if year % 4 == 0 or (year % 400 == 0 and year % 100 != 0):
        return True
    return False


# given the first day Mon 1/1/1900 determine the first day of the given year
# based on total number of days elapsed
def first_day_of_year(day_one, weekdays, year, years):
    day_gap = 0
    for y in years:
        if y["year"] >= year:
            break
        else:
            day_gap += y["days"]
    
    print(day_gap, weekdays[day_gap%len(weekdays)-1])
    return weekdays[day_gap%len(weekdays)-1], day_gap


# determine the days (of the week) of the month for the given year
def days_of_the_month(day_one, weekdays, years):
    for y in years:
        first_day_of_yr, day_gap = first_day_of_year(day_one, weekdays, y["year"], years)
        y["first_day"] = first_day_of_yr
        # set 1/1/year to first_day
        y["months"][0]["first_day"] = first_day_of_yr
        first_day_offset = weekdays.index(first_day_of_yr)

        for m in y["months"]:
            day_gap += m["days"]
            if "first_day" not in m:
                m["first_day"] = weekdays[day_gap%(len(weekdays)-1)]
    return years
        

# one way to tackle this without overexerting ourselves is to focus on 
# determining the first day of each month for every year in the range
# which is simply:
#   1/month/year
#
# given we know the number of days in a month (and year) it becomes a 
# simpler math problem
#
#
#
def main(day_one, start_range, end_range):
    yrs = gen_years(start_range["year"], end_range["year"])
    
    weekdays = gen_weekdays()

    # first_day_of_year(day_one, weekdays, 1901, yrs)
    days_of_the_month(day_one, weekdays, yrs)
    pprint.pprint(yrs)


if __name__ == "__main__":
    day_one = {"year": 1900, "month": "jan", "day": "mon", "day_of_month": 1}
    start_range = {"year": 1901, "month": "jan", "day_of_month": 1}
    end_range = {"year": 2000, "month": "dec", "day_of_month": 31}

    main(day_one, start_range, end_range)