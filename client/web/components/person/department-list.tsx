import type { Credits, MovieCast, MovieCrew, TVCast, TVCrew } from '@/models/people-types';
import { ChevronDown } from 'lucide-react';
import Department from './department';
import { useState } from 'react';
// import { Button } from '@/components/ui/button';

interface DepartmentListProps {
  credits: Credits;
  knownForDepartment: string;
}

function DropdownItem({ label, number, onClick }: { label: string; number: number; onClick: () => void }) {
  return (
    <div
      className="flex justify-between px-4 py-2 text-sm rounded-lg hover:bg-gray-200 hover:text-gray-900"
      onClick={onClick}
    >
      <div>{label}</div>
      <div className="text-gray-400">{number}</div>
    </div>
  );
}

function getDepartmentArray(credits: Credits): { name: string; count: number }[] {
  const departmentCounts: Record<string, number> = {};

  // Lấy department từ crew
  [...credits.movie_credits.crew, ...credits.tv_credits.crew].forEach((crewMember) => {
    departmentCounts[crewMember.department] = (departmentCounts[crewMember.department] || 0) + 1;
  });

  // Thêm "Acting" nếu có cast
  const actingCount = credits.movie_credits.cast.length + credits.tv_credits.cast.length;
  if (actingCount > 0) {
    departmentCounts['Acting'] = actingCount;
  }

  // Chuyển đổi từ Record sang mảng { name, count }
  return Object.entries(departmentCounts).map(([name, count]) => ({ name, count }));
}

function groupAndSortByDepartment(
  credits: Credits,
  knownForDepartment: string,
  options?: {
    type?: 'movie' | 'tv'; // Lọc theo thể loại movie hoặc tv
    filterDepartment?: string; // Lọc theo một department cụ thể
  },
): { department: string; items: (MovieCrew | TVCrew | MovieCast | TVCast)[] }[] {
  const departmentGroups: Record<string, (MovieCrew | TVCrew | MovieCast | TVCast)[]> = {};

  // Lấy dữ liệu cast và crew dựa trên type
  const castData =
    options?.type === 'movie'
      ? credits.movie_credits.cast
      : options?.type === 'tv'
      ? credits.tv_credits.cast
      : [...credits.movie_credits.cast, ...credits.tv_credits.cast];

  const crewData =
    options?.type === 'movie'
      ? credits.movie_credits.crew
      : options?.type === 'tv'
      ? credits.tv_credits.crew
      : [...credits.movie_credits.crew, ...credits.tv_credits.crew];

  // Gom nhóm cast thành nhóm Acting
  if (castData.length > 0) {
    departmentGroups['Acting'] = castData;
  }

  // Gom nhóm crew theo department
  crewData.forEach((crewMember) => {
    if (!departmentGroups[crewMember.department]) {
      departmentGroups[crewMember.department] = [];
    }
    departmentGroups[crewMember.department].push(crewMember);
  });

  // Lọc theo department nếu cần
  if (options?.filterDepartment) {
    Object.keys(departmentGroups).forEach((department) => {
      if (department !== options.filterDepartment) {
        delete departmentGroups[department];
      }
    });
  }

  // Sắp xếp các phần tử trong từng nhóm theo năm (giảm dần)
  Object.keys(departmentGroups).forEach((department) => {
    departmentGroups[department].sort((a, b) => {
      const yearA = getYear('release_date' in a ? a.release_date : a.first_air_date);
      const yearB = getYear('release_date' in b ? b.release_date : b.first_air_date);
      return yearB - yearA; // Sắp xếp giảm dần theo năm
    });
  });

  // Chuyển thành mảng { department, items } và sắp xếp các nhóm
  const sortedGroups = Object.entries(departmentGroups)
    .map(([department, items]) => ({ department, items }))
    .sort((a, b) => {
      // Đưa nhóm knownForDepartment lên đầu
      if (a.department === knownForDepartment) return -1;
      if (b.department === knownForDepartment) return 1;

      // Sắp xếp theo số lượng phần tử trong nhóm (giảm dần)
      return b.items.length - a.items.length;
    });

  return sortedGroups;
}

function getYear(dateString: string | undefined): number {
  return dateString ? parseInt(dateString.split('-')[0], 10) : 0;
}

const DepartmentList: React.FC<DepartmentListProps> = ({ credits, knownForDepartment }) => {
  const [selectedType, setSelectedType] = useState<'movie' | 'tv' | undefined>(undefined);
  const [selectedDepartment, setSelectedDepartment] = useState<string | undefined>(undefined);

  const handleTypeSelect = (type: 'movie' | 'tv' | undefined) => {
    setSelectedType(type);
  };

  const handleDepartmentSelect = (department: string | undefined) => {
    setSelectedDepartment(department);
  };

  return (
    <div className="pb-10 px-5">
      <div className="container mx-auto text-black">
        <div className="flex flex-col justify-between items-start ml-10 gap-5">
          <div className="flex items-center justify-end w-full">
            <div className="flex space-x-5 text-black text-lg font-medium gap-5">
              <div className="relative group hover:cursor-pointer">
                {/* <Button */}
                <span className="flex items-center gap-2">
                  All <ChevronDown />
                </span>
                <div className="absolute left-0 -translate-x-14 hidden w-48 bg-white text-gray-800 rounded-lg shadow-lg group-hover:block">
                  <DropdownItem
                    number={credits.movie_credits.cast.length + credits.movie_credits.crew.length}
                    label="Movies"
                    onClick={() => handleTypeSelect('movie')}
                  />
                  <DropdownItem
                    number={credits.tv_credits.cast.length + credits.tv_credits.crew.length}
                    label="TV Shows"
                    onClick={() => handleTypeSelect('tv')}
                  />
                </div>
              </div>
              <div className="relative group hover:cursor-pointer">
                <span className="flex items-center gap-2">
                  Department <ChevronDown />
                </span>
                <div className="absolute left-0 -translate-x-14 hidden w-48 bg-white text-gray-800 rounded-lg shadow-lg group-hover:block">
                  {getDepartmentArray(credits).map((department, index) => (
                    <DropdownItem
                      key={index}
                      number={department.count}
                      label={department.name}
                      onClick={() => handleDepartmentSelect(department.name)}
                    />
                  ))}
                </div>
              </div>
            </div>
          </div>
          {groupAndSortByDepartment(credits, knownForDepartment, {
            type: selectedType,
            filterDepartment: selectedDepartment,
          }).map((element, index) => (
            <Department key={index} props={element} />
          ))}
        </div>
      </div>
    </div>
  );
};

export default DepartmentList;
