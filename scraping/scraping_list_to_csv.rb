# frozen_string_literal: true

require 'httparty'
require 'nokogiri'
require 'csv'

response = HTTParty.get('http://www.hkt48.jp/profile/')
body = response.body
html_doc = Nokogiri::HTML5(body)

team_info = {}
teams = ['Team H', 'Team KIV', 'Team TII']

html_doc.css('.profile_list ul').each_with_index do |team, index|
  num_member = team.css('li').length
  team_info[teams[index]] = num_member
end

def get_team_name(index, team_info)
  counts = team_info.values
  additional_counts = []
  counts.each_with_index do |count, i|
    if i.zero?
      additional_counts << count
    else
      additional_counts.push(additional_counts[i - 1] + count)
    end
  end

  additional_counts.each.with_index do |count, i|
    next if index + 1 > count

    return team_info.keys[i]
  end
end

CSV.open('team.csv', 'w') do |csv|
  csv << %w[team_name num_member]
  team_info.each do |team|
    csv << team.to_a
  end
end

CSV.open('member.csv', 'w') do |csv|
  csv << %w[num name profile_link team_name]
  html_doc.css('.profile_list ul li a').each_with_index do |member, index|
    csv << [index + 1, member.text, "http://www.hkt48.jp#{member['href']}", get_team_name(index, team_info)]
  end
end
