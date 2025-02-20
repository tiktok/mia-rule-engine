/*
 * Copyright 2022 TikTok Pte. Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package benchmark

const POLICY = `
DEFAULT GRANT $scope AS account:normal; 
ADD $fact AS $subject.age_status:underage WHEN $subject.age < 14 AND ($subject.geo = "KR" OR $subject.geo = "ID");
ADD $fact AS $subject.age_status:underage WHEN $subject.age < 13 AND $subject.geo = "US";
ADD $fact AS $subject.age_status:adult WHEN $subject.age >= 18;
GRANT $scope AS account:kids_mode WHEN $subject.age_status = "underage" AND $subject.geo = "US" PRIORITY 1;
GRANT $scope AS account:banned WHEN $subject.age_status = "underage" PRIORITY 2;
`
